package sources

import (
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/stream"
)

type FaninSource struct {
	rawChan chan stream.IMessage
	sources []RawSource
}

func NewFaninSourceWithKafka(brokers, topic, group string, partitions int, opts ...KafkaOption) (RawSource, error) {
	rawChan := make(chan stream.IMessage, partitions)
	sources := make([]RawSource, partitions)
	if opts == nil {
		opts = make([]KafkaOption, 0, 1)
	}
	opts = append(opts, WithRawChan(rawChan))
	zap.L().Info("NewFaninSourceWithKafka!!!", zap.String("topic", topic), zap.String("group", group))
	for i := 0; i < partitions; i++ {
		src, err := NewKafkaSource(brokers, topic, group, opts...)
		if err != nil {
			return nil, err
		}
		sources[i] = src
	}
	return &FaninSource{
		rawChan: rawChan,
		sources: sources,
	}, nil
}

func NewFaninSourceWithNATs(server, subject, queue string, workers int, opts ...NATsOption) (RawSource, error) {
	rawChan := make(chan stream.IMessage, workers)
	sources := make([]RawSource, workers)
	if opts == nil {
		opts = make([]NATsOption, 0, 1)
	}
	opts = append(opts, WithRawChanNATs(rawChan))
	// zap.L().Info("NewFaninSourceWithKafka!!!", zap.String("topic", topic), zap.String("group", group))
	for i := 0; i < workers; i++ {
		src, err := NewNATsSource(server, subject, queue, opts...)
		if err != nil {
			return nil, err
		}
		sources[i] = src
	}
	return &FaninSource{
		rawChan: rawChan,
		sources: sources,
	}, nil
}

func (s *FaninSource) RawBytes() chan stream.IMessage {
	return s.rawChan
}

func (s *FaninSource) Start() {
	for _, src := range s.sources {
		go src.Start()
	}
}

func (s *FaninSource) Close() {
	for _, src := range s.sources {
		src.Close()
	}
}
