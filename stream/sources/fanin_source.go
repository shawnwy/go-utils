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
