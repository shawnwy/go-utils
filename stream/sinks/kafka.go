package sinks

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/errors"
	"github.com/shawnwy/go-utils/v5/stream"
)

type KafkaSink struct {
	topic string
	p     *kafka.Producer
	stop  chan struct{}
	once  sync.Once
	wg    sync.WaitGroup
	cfg   *kafka.ConfigMap
}

type KafkaOption func(sink *KafkaSink)

func WithClientID(clientID string) KafkaOption {
	return func(s *KafkaSink) {
		_ = s.cfg.SetKey("client.id", clientID)
	}
}

func WithACKs(ack int) KafkaOption {
	return func(s *KafkaSink) {
		_ = s.cfg.SetKey("acks", ack)
	}
}

func WithCompression(ctype string, lv int) KafkaOption {
	return func(s *KafkaSink) {
		_ = s.cfg.SetKey("compression.type", ctype)
		_ = s.cfg.SetKey("compression.level", lv)
	}
}

func WithConfigMap(cfg kafka.ConfigMap) KafkaOption {
	return func(s *KafkaSink) {
		for k, v := range cfg {
			_ = s.cfg.SetKey(k, v)
		}
	}
}

func NewKafkaSink(brokers, topic string, opts ...KafkaOption) (Sink, error) {
	s := &KafkaSink{
		stop:  make(chan struct{}),
		topic: topic,
		cfg: &kafka.ConfigMap{
			"bootstrap.servers":   brokers,
			"client.id":           fmt.Sprintf("infra-sink-%d", rand.Intn(1000)),
			"go.batch.producer":   true,
			"linger.ms":           1000,
			"compression.type":    "lz4",
			"compression.level":   1,
			"batch.num.messages":  10000,
			"acks":                1,
			"go.delivery.reports": false,
		},
	}
	for _, o := range opts {
		o(s)
	}
	var err error
	if s.p, err = kafka.NewProducer(s.cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to construct Kafka Producer")
	}
	return s, nil
}

func (k *KafkaSink) Subscribe(ingress <-chan stream.IMessage) {
	k.wg.Add(1)
	defer k.wg.Done()

	for {
		select {
		case m, ok := <-ingress:
			if !ok {
				return
			}

			if err := k.p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &k.topic,
					Partition: kafka.PartitionAny,
				},
				Value: m.Bytes(),
				Key:   m.UID(),
			}, nil); err != nil {
				zap.L().Info("failed to sink a msg from kafka", zap.Error(err))
			}
		case <-k.stop:
			zap.L().Info("kafka sink has been stopped")
			return
		}
	}
}

func (k *KafkaSink) Wait() {
	k.wg.Wait()
}

func (k *KafkaSink) Close() {
	k.once.Do(func() {
		close(k.stop)
		k.p.Flush(15 * 1000)
		k.p.Close()
	})
}

func (k *KafkaSink) HandleError(cb func(err interface{})) {
	if cb == nil {
		cb = defaultErrCB
	}
	for event := range k.p.Events() {
		switch event.(type) {
		case kafka.Error:
			cb(nil)
		}
	}
}
