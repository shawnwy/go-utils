package sinks

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/gen"
	"github.com/shawnwy/go-utils/v5/stream"
)

func TestNATsSink(t *testing.T) {
	dev, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(dev)
	sink, err := NewNATsSink("localhost:4222", "test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ingress := make(chan stream.IMessage, 10)
	m := make(stream.RawMessage, 10)
	gen.RandomBytesLite(m)
	go sink.Subscribe(ingress)
	go sink.HandleError(func(err interface{}) {
		fmt.Println(err)
		fmt.Println("1")
	})
	for i := 0; i < 10; i++ {
		ingress <- m
	}
	time.Sleep(10 * time.Second)
}
