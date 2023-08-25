package sources

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/errors"
	"github.com/shawnwy/go-utils/v5/stream"
)

type NATsSource struct {
	rawChan chan stream.IMessage
	subConn *nats.Conn
	sub     *nats.Subscription

	stop chan struct{}
	once sync.Once
}

type NATsOption func(source *NATsSource)

// WithRawChanNATs set rawChan for NATsSource
func WithRawChanNATs(rawChan chan stream.IMessage) NATsOption {
	return func(s *NATsSource) {
		s.rawChan = rawChan
	}
}

func NewNATsSource(server, subject, queue string, opts ...nats.Option) (_ RawSource, err error) {
	s := &NATsSource{
		stop: make(chan struct{}),
	}
	opts = append([]nats.Option{
		nats.MaxReconnects(-1),
		nats.RetryOnFailedConnect(true),
		nats.ReconnectJitter(time.Second, time.Second),
		nats.ReconnectHandler(func(c *nats.Conn) {
			zap.L().Info("nats connection has been established", zap.Uint64("retries#", c.Reconnects))
		}),
		nats.Name(fmt.Sprintf("src.%s", subject)),
	}, opts...)
	s.subConn, err = nats.Connect(server, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make a NATs connection")
	}
	// initialize raw chan
	if s.rawChan == nil {
		s.rawChan = make(chan stream.IMessage, rawSrcChanSize)
	}
	s.sub, err = s.subConn.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		s.rawChan <- stream.RawMessage(msg.Data)
	})
	return s, nil
}

func (s *NATsSource) RawBytes() chan stream.IMessage {
	return s.rawChan
}

func (s *NATsSource) Start() {
	defer s.subConn.Close()
	for {
		select {
		case <-s.stop:
			err := s.sub.Unsubscribe()
			if err != nil {
				zap.L().Warn("failed to unsubscribe", zap.Error(err))
			}
			return
		}
	}
}

func (s *NATsSource) Close() {
	s.once.Do(func() {
		close(s.stop)
	})
}
