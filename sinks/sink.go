package sinks

type Sink interface {
	Subscribe(ingress <-chan []byte)
	HandleError(cb func(err interface{}))
	Wait()
	Close()
}
