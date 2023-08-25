package rotatelogs

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"

	"github.com/shawnwy/go-utils/v5/errors"
)

var (
	PcapExt     string = ".pcap"
	MsgBytesExt string = ".mb"

	NewLine            = []byte("\r\n|\r\n")
	ErrInvalidFn error = errors.New("err: invalid filename")
)

// LogFile is a struct holding info of log files. Its ext could be *.pcap/*.mb etc.
// 	- Node: where the packets capture for pcap file,
//			usually is the host IP
//	- Timestamp: when the pcap file generated
//	- Filepath: the file path of  pcap file
type LogFile struct {
	Node      string
	Timestamp time.Time
	Filepath  string
}

func NewLogFile(abspath, ext string, fnparser FilenameParser) (f LogFile, err error) {
	filename := strings.TrimSuffix(filepath.Base(abspath), ext)
	node, timestamp, err := fnparser(filename)
	if err != nil {
		err = errors.Wrap(err, "failed to scan files")
		return
	}
	return LogFile{
		Filepath:  abspath,
		Node:      node,
		Timestamp: timestamp,
	}, nil
}

func (f LogFile) AsFile() (*os.File, error) {
	return os.Open(f.Filepath)
}

// Shred - physically remove the file from os
func (f LogFile) Shred() error {
	return os.Remove(f.Filepath)
}

// HeapQxMap is a map with many priority Queue sorted by key
// and the queue in map is not thread-safe
type HeapQxMap struct {
	M       map[string]*priorityqueue.Queue
	orderBy func(a, b interface{}) int
}

func New(orderF func(a, b interface{}) int) *HeapQxMap {
	return &HeapQxMap{
		M:       make(map[string]*priorityqueue.Queue),
		orderBy: orderF,
	}
}

func (h *HeapQxMap) Enqueue(f LogFile) {
	q, ok := h.M[f.Node]
	if !ok {
		q = priorityqueue.NewWith(h.orderBy)
		h.M[f.Node] = q
	}
	q.Enqueue(f)
}

func (h *HeapQxMap) GetQueue(key string) *priorityqueue.Queue {
	q, ok := h.M[key]
	if !ok {
		q = priorityqueue.NewWith(h.orderBy)
		h.M[key] = q
	}
	return q
}

// ByTime time natural ordering
func ByTime(a, b interface{}) int {
	pre := a.(LogFile)
	next := b.(LogFile)

	return utils.TimeComparator( // Note "-" for descending order
		pre.Timestamp,
		next.Timestamp,
	)
}

// SplitAt Custom split function. This will split string at 'token' i.e # or // etc....
func SplitAt(tk []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchLen := len(tk)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return Nothing if at the end of file or no data passed.
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token.
		if i := bytes.Index(data, tk); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}
