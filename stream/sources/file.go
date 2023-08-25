package sources

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/errors"
	"github.com/shawnwy/go-utils/v5/rotatelogs"
	"github.com/shawnwy/go-utils/v5/stream"
)

const (
	defaultWorkers = 1
	defaultSnapMS  = 200
)

type FileSourceOption func(source *FileSource)

// WithFnameParser - set the filename parser for FileSource
func WithFnameParser(p rotatelogs.FilenameParser) FileSourceOption {
	return func(s *FileSource) {
		s.fnParser = p
	}
}

// WithSnapMS - set the snapMS for FileSource, which indicate
// how long to wait for next scan if no files after the scanned.
func WithSnapMS(d int64) FileSourceOption {
	return func(s *FileSource) {
		s.snapMS = d
	}
}

// WithWorkers - set the number of parallel Queue processing
func WithWorkers(n int) FileSourceOption {
	return func(s *FileSource) {
		s.workers = n
	}
}

// WithShred - active shredder of FileSource. files will not keep after processed.
func WithShred() FileSourceOption {
	return func(s *FileSource) {
		s.shred = true
	}
}

// WithFileProcessorFactory - set ProcessorFactory for source
func WithFileProcessorFactory(f ProcessorFactory) FileSourceOption {
	return func(s *FileSource) {
		s.factory = f
	}
}

// AddFileValidator - add a new FileValidator for source
func AddFileValidator(v FileValidator) FileSourceOption {
	return func(s *FileSource) {
		s.validators = append(s.validators, v)
	}
}

// FileValidator validate the LogFile whether is suitable for processing
type FileValidator func(f rotatelogs.LogFile, now time.Time) error

type FileSource struct {
	workers  int
	dir      string
	ext      string
	snapMS   int64
	m        *rotatelogs.HeapQxMap
	fnParser rotatelogs.FilenameParser
	rawChan  chan stream.IMessage
	stop     func()
	shred    bool // whether to keep files

	factory    ProcessorFactory
	validators []FileValidator
}

func NewFileSource(
	dir, ext string,
	fnParser rotatelogs.FilenameParser,
	opts ...FileSourceOption,
) (_ RawSource, err error) {
	s := &FileSource{
		workers:    defaultWorkers,
		dir:        dir,
		ext:        ext,
		snapMS:     defaultSnapMS,
		m:          rotatelogs.New(rotatelogs.ByTime),
		fnParser:   fnParser,
		factory:    defaultProcessorFactory(),
		validators: make([]FileValidator, 0),
	}
	for _, o := range opts {
		o(s)
	}

	if s.rawChan == nil {
		s.rawChan = make(chan stream.IMessage, rawSrcChanSize)
	}
	return s, nil
}

func (s *FileSource) RawBytes() chan stream.IMessage {
	return s.rawChan
}

func (s *FileSource) Close() {
	s.stop()
}

func (s *FileSource) Start() {
	s.watch(context.Background())
}

// watch spawn files in the directory and continuous processes them
func (s *FileSource) watch(ctx context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	// start workers for parallel process
	queueCh := make(chan *priorityqueue.Queue, s.workers)
	wg := sync.WaitGroup{}
	for i := 0; i < s.workers; i++ {
		go s.processQ(childCtx, queueCh, &wg)
	}

	for {
		select {
		case <-ctx.Done():
			zap.L().Info("FileSource is stopping watch ...")
			return

		default:
			n, err := s.scan()
			if err != nil {
				zap.L().Warn("failed to scan directory", zap.Error(err))
				continue
			}
			if n == 0 {
				// snap a while for next scan
				zap.L().Info("NONE file spawn ...")
				time.Sleep(time.Duration(s.snapMS) * time.Millisecond)
				continue
			}
			// Process files
			wg.Add(len(s.m.M))
			for _, q := range s.m.M {
				queueCh <- q
			}
			wg.Wait() // wait for all workers processed all queues
			time.Sleep(time.Second * 30)
		}
	}
}

// scan the directory and discover spawned files
func (s *FileSource) scan() (n int, err error) {
	scanT := time.Now()
	files, err := filepath.Glob(filepath.Join(s.dir, "*"+s.ext))
	if err != nil {
		err = errors.Wrap(err, "failed to scan files")
		return
	}
	n = len(files)
	if n > 0 {
		zap.L().Info("start scan files ...")
	}

	var queuedCnt int
	for _, f := range files {
		lf, err := rotatelogs.NewLogFile(f, s.ext, s.fnParser)
		if err != nil {
			zap.L().Warn("failed to create logfile", zap.Error(err), zap.String("fname", f))
			continue
		}

		for _, v := range s.validators {
			if err := v(lf, scanT); err != nil {
				zap.L().Warn("failed to validate LogFile", zap.Error(err))
				continue
			}
		}

		s.m.Enqueue(lf)
		queuedCnt++
	}
	zap.L().Info("scan has been done ...",
		zap.Int("queued #files", queuedCnt),
		zap.Int("scanned #files", n))

	return n, nil
}

// processQ is worker goroutine for processing a file queue
func (s *FileSource) processQ(ctx context.Context, qCh chan *priorityqueue.Queue, wg *sync.WaitGroup) {
	processor := s.factory(s.rawChan)

	for {
		select {
		case <-ctx.Done():
			zap.L().Info("stopping Process fileQ ...")
			return

		case q := <-qCh:
			// drain off the LogFile Q

			// preprocess func
			for !q.Empty() {
				item, ok := q.Dequeue()
				if !ok {
					zap.L().Warn("failed to pop logfile Q")
					continue
				}
				lf, ok := item.(rotatelogs.LogFile)
				if !ok {
					zap.L().Warn("failed to type the interface as LogFile")
					continue
				}
				// open the LogFile
				file, err := lf.AsFile()
				if err != nil {
					zap.L().Warn("failed to Process the file",
						zap.String("fname", lf.Filepath),
						zap.Error(err))
					return
				}
				processor.Process(file)
				// close file after Process
				if err = file.Close(); err != nil {
					zap.L().Warn("failed to close file",
						zap.String("fname", lf.Filepath),
						zap.Error(err))
				}

				if !s.shred {
					continue // no need to shred file, skip
				}
				// shred the processed file
				if err := lf.Shred(); err != nil {
					zap.L().Warn("failed to shred file after readout",
						zap.String("fname", lf.Filepath),
						zap.Error(err))
					continue
				}
			}
			wg.Done()
		}
	}
}

// ProcessorFactory - is used by FileSource to create LogProcessor
type ProcessorFactory func(chan stream.IMessage) LogProcessor

// LogProcessor - help to Process the LogFile
type LogProcessor interface {
	Process(file *os.File) // process logfile
}

// DefaultProcessor - is a default LogProcessor read out contents directly from file to rawChan channel
type DefaultProcessor struct {
	r      *bufio.Reader
	egress chan stream.IMessage
}

// defaultProcessFactory - will generate a factory create default processor
func defaultProcessorFactory() ProcessorFactory {
	return func(egress chan stream.IMessage) LogProcessor {
		return &DefaultProcessor{
			r:      bufio.NewReader(nil),
			egress: egress,
		}
	}
}

func (d *DefaultProcessor) Process(file *os.File) {
	d.r.Reset(file)
	sc := bufio.NewScanner(d.r)
	sc.Split(rotatelogs.SplitAt(rotatelogs.NewLine))
	for sc.Scan() {
		d.egress <- stream.RawMessage(sc.Bytes())
	}
	if err := sc.Err(); err != nil {
		zap.L().Warn("failed to scan a file", zap.Error(err))
	}
}
