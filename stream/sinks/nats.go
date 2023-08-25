package sinks

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/stream"
)

type NATsSink struct {
	subject string
	pubConn *nats.Conn
	errCh   chan error

	stop chan struct{}
	once sync.Once
	wg   sync.WaitGroup
}

func NewNATsSink(server, subject string, opts ...nats.Option) (_ Sink, err error) {
	s := &NATsSink{
		subject: subject,
		errCh:   make(chan error, 10),

		stop: make(chan struct{}),
	}
	opts = append([]nats.Option{
		nats.MaxReconnects(-1),
		nats.RetryOnFailedConnect(true),
		nats.ReconnectJitter(time.Second, time.Second),
		nats.ReconnectHandler(func(c *nats.Conn) {
			zap.L().Info("nats connection has been established", zap.Uint64("retries#", c.Reconnects))
		}),
		nats.Name("nats-sink"),
		// nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
		// 	s.errCh <- err
		// }),
	}, opts...)
	s.pubConn, err = nats.Connect(server, opts...)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *NATsSink) Subscribe(ingress <-chan stream.IMessage) {
	s.wg.Add(1)
	defer s.wg.Done()
	for {
		select {
		case m := <-ingress:
			bytes := m.Bytes()
			if err := s.pubConn.Publish(s.subject, bytes); err != nil {
				zap.L().Warn("failed to pub through nats", zap.Error(err), zap.Int("size", len(bytes)))
				fmt.Println("1")
				continue
			}

		case <-s.stop:
			zap.L().Info("nats sink has been stopped")
			return
		}
	}
}

func (s *NATsSink) HandleError(cb func(err interface{})) {
	if cb == nil {
		cb = defaultErrCB
	}
	for e := range s.errCh {
		cb(e)
	}
}

func (s *NATsSink) Wait() {
	s.wg.Wait()
}

func (s *NATsSink) Close() {
	s.once.Do(func() {
		close(s.stop)
		if err := s.pubConn.Drain(); err != nil {
			zap.L().Warn("failed to drain and close conn", zap.Error(err))
			return
		}
	})
}
