package sinks

import (
	"github.com/shawnwy/go-utils/v5/errors"
)

type FanoutSink struct {
	sinks []Sink
}

func (s *FanoutSink) HandleError(cb func(err interface{})) {
	for _, sk := range s.sinks {
		go sk.HandleError(cb)
	}
}

func NewFanoutSinkWithKafka(brokers, topic string, partitions int, opts ...KafkaOption) (Sink, error) {
	sinks := make([]Sink, partitions)
	for i := 0; i < partitions; i++ {
		sink, err := NewKafkaSink(brokers, topic, opts...)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create fan-out sink")
		}
		sinks[i] = sink
	}
	return &FanoutSink{sinks}, nil
}

func NewFanoutSinkWithSyslog(proto, url, tag string, workers int) (Sink, error) {
	sinks := make([]Sink, workers)
	for i := 0; i < workers; i++ {
		sink, err := NewSyslogSink(proto, url, tag)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create fan-out sink")
		}
		sinks[i] = sink
	}
	return &FanoutSink{sinks}, nil
}

func (s *FanoutSink) Subscribe(ingress <-chan []byte) {
	for _, sk := range s.sinks {
		go sk.Subscribe(ingress)
	}
}

func (s *FanoutSink) Wait() {
	for _, sk := range s.sinks {
		sk.Wait()
	}
}

func (s *FanoutSink) Close() {
	for _, sk := range s.sinks {
		sk.Close()
	}
}
