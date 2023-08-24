package sinks

import (
	"log/syslog"
	"strings"
	"sync"

	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/errors"
	"github.com/shawnwy/go-utils/v5/stream"
)

type SyslogSink struct {
	tag string
	wr  *syslog.Writer

	stop  chan struct{}
	once  sync.Once
	wg    sync.WaitGroup
	errCh chan struct{}
}

func NewSyslogSink(proto, url, tag string) (Sink, error) {
	wr, err := syslog.Dial(strings.ToLower(proto), url, syslog.LOG_INFO, tag)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create syslog sink")
	}
	return &SyslogSink{
		wr:    wr,
		tag:   tag,
		stop:  make(chan struct{}),
		errCh: make(chan struct{}),
	}, nil
}

func (s *SyslogSink) Subscribe(ingress <-chan stream.IMessage) {
	s.wg.Add(1)
	defer s.wg.Done()
	for {
		select {
		case m, ok := <-ingress:
			if !ok {
				zap.L().Info("syslog sink.exit: ingress channel has been closed")
				return
			}
			bytes := m.Bytes()
			if n, err := s.wr.Write(bytes); err != nil || n < len(bytes) {
				if err != nil {
					zap.L().Warn("failed to sink a msg from syslog", zap.Error(err))
					s.errCh <- struct{}{}
					return
				}

				zap.L().Warn("truncated")
			}

		case <-s.stop:
			zap.L().Info("syslog sink has been stopped")
			return
		}
	}
}

func (s *SyslogSink) Wait() {
	s.wg.Wait()
}

func (s *SyslogSink) Close() {
	s.once.Do(func() {
		close(s.stop)
		err := s.wr.Close()
		if err != nil {
			zap.L().Warn("failed to closed syslog writer", zap.Error(err))
			return
		}
		zap.L().Info("syslog sink has been closed successfully")
	})
}

func (s *SyslogSink) HandleError(cb func(err interface{})) {
	for range s.errCh {
		cb(nil)
	}
}
