package sources

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewKafkaSource(t *testing.T) {
	src, err := NewKafkaSource("10.10.40.115:9092", "log-data-topic", "test-close-kafka")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	go func() {
		var cnt int64
		for range src.RawBytes() {
			cnt++
			if cnt%10000 == 0 {
				println("comsumed #", cnt)
			}
		}
	}()
	go src.Start()
	time.Sleep(time.Second * 5)
	println("closing ...")
	src.Close()
	println("closed")
	time.Sleep(time.Second * 5)
}

func F(c context.Context, id string) {
	<-c.Done()
	fmt.Printf("F#%s has been done\n", id)
}

func TestContext(t *testing.T) {
	c, cancel := context.WithCancel(context.Background())

	go F(c, "1")
	go F(c, "2")
	go F(c, "3")
	cancel()
	time.Sleep(5 * time.Second)
}

func TestContextPropagate(t *testing.T) {
	c1, cancel := context.WithCancel(context.Background())
	
	c2, _ := context.WithCancel(c1)
	// defer cancel2()
	c3, cancel3 := context.WithCancel(c2)
	// defer cancel3()
	go F(c1, "1")
	go F(c2, "2")
	go F(c3, "3")
	cancel3()
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
