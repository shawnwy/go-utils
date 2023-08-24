package sources

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/shawnwy/go-utils/v5/errors"
	"github.com/shawnwy/go-utils/v5/sizes"
	"github.com/shawnwy/go-utils/v5/stream"

	"go.uber.org/zap"
)

type KafkaSource struct {
	clientID string
	rawChan  chan stream.IMessage
	c        *kafka.Consumer
	stop     chan struct{}
	once     sync.Once
}

type KafkaOption func(source *KafkaSource)

// WithRawChan set rawChan for KafkaSource
func WithRawChan(rawChan chan stream.IMessage) KafkaOption {
	return func(s *KafkaSource) {
		s.rawChan = rawChan
	}
}

// WithClientID set clientID for KafkaSource
func WithClientID(clientID string) KafkaOption {
	return func(s *KafkaSource) {
		s.clientID = clientID
	}
}

func NewKafkaSource(brokers, topic, group string, opts ...KafkaOption) (RawSource, error) {
	s := &KafkaSource{
		stop:     make(chan struct{}),
		clientID: fmt.Sprintf("infra-source-%d", rand.Intn(1000)),
	}
	for _, o := range opts {
		o(s)
	}

	var err error
	s.c, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		// Use the cooperative incremental rebalance protocol.
		"go.application.rebalance.enable": true,
		// "partition.assignment.strategy":   "cooperative-sticky",
		"group.id":           group,
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "latest",
		// "enable.auto.offset.store":        false,
		"fetch.message.max.bytes":    8 * int(sizes.MB),
		"queued.max.messages.kbytes": 64 * int(sizes.MB/sizes.KB),
		"client.id":                  s.clientID,
	})

	if err != nil {
		zap.Error(err)
		return nil, errors.Wrap(err, "failed to construct Kafka Consumer")
	}

	if s.rawChan == nil {
		s.rawChan = make(chan stream.IMessage, rawSrcChanSize)
	}

	err = s.c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe topic")
	}
	return s, nil
}

func (k *KafkaSource) RawBytes() chan stream.IMessage {
	if k.rawChan == nil {
		k.rawChan = make(chan stream.IMessage, rawSrcChanSize)
	}
	return k.rawChan
}

// Start activate kafka source
//
//	caution: RawBytes() has to be called before Start()
func (k *KafkaSource) Start() {
	defer k.c.Close()
	for {
		select {
		case <-k.stop:
			return

		default:
			event := k.c.Poll(100)
			if event == nil {
				continue
			}

			switch e := event.(type) {
			case *kafka.Message:
				if e.TopicPartition.Error != nil {
					zap.L().Warn("kafka topic partition error raised", zap.Error(e.TopicPartition.Error))
					continue
				}
				k.rawChan <- stream.RawMessage(e.Value)
			case kafka.Error:
				zap.L().Warn("kafka error raised", zap.Error(e))
			}
		}
	}
}

func (k *KafkaSource) Close() {
	k.once.Do(func() {
		close(k.stop)
	})
}
