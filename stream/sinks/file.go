package sinks

import (
	"io"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/shawnwy/go-utils/v5/pathutils"
	"github.com/shawnwy/go-utils/v5/rotatelogs"
	"github.com/shawnwy/go-utils/v5/sizes"
	"github.com/shawnwy/go-utils/v5/stream"
)

type FileEmitter struct {
	w io.Writer

	stop chan struct{}
	once sync.Once
}

func NewFileEmitter(path, fname string) *FileEmitter {
	wr := &lumberjack.Logger{
		Filename: filepath.Join(pathutils.AbsPath(path), fname),
		MaxSize:  1 * int(sizes.GB), // megabytes
		MaxAge:   28,                // days
		Compress: false,             // disabled by default
	}
	return &FileEmitter{w: wr}
}

func (e *FileEmitter) Subscribe(ingress <-chan stream.IMessage) {
	for {
		select {
		case m := <-ingress:
			_, err := e.w.Write(m.Bytes())
			if err != nil {
				zap.L().Warn("failed to emit to file", zap.Error(err))
			}
			if _, err = e.w.Write(rotatelogs.NewLine); err != nil {
				zap.L().Warn("failed to emit to NewLine", zap.Error(err))
			}

		case <-e.stop:
			zap.L().Info("stopping emit ...")
			return
		}
	}
}

func (e *FileEmitter) HandleError(cb func(err interface{})) {

}

func (e *FileEmitter) Wait() {

}

func (e *FileEmitter) Close() {
	e.once.Do(func() {
		close(e.stop)
	})
}
