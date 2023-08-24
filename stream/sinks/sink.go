package sinks

import "github.com/shawnwy/go-utils/v5/stream"

type Sink interface {
	Subscribe(ingress <-chan stream.IMessage)
	HandleError(cb func(err interface{}))
	Wait()
	Close()
}

var defaultErrCB = func(interface{}) {}
