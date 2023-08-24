package sinks

import (
	"fmt"
	"log/syslog"
	"testing"

	"github.com/shawnwy/go-utils/v5/gen"
	"github.com/shawnwy/go-utils/v5/sizes"
)

var (
	benchWriter, _ = syslog.Dial("udp", "10.10.20.68:514", syslog.LOG_INFO, "test-bench")

	payload = make([]byte, 10*sizes.KB)
)

func TestLog(t *testing.T) {

	logWriter, err := syslog.Dial("udp", "10.10.20.68:514", syslog.LOG_INFO, "infra-pipe")
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	err = logWriter.Info("test1")
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
}

func TestSyslogSink_Sink(t *testing.T) {
	sink, err := NewSyslogSink("udp", "10.10.20.68:514", "infra-pipe-syslog-test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ingress := make(chan []byte, 100)
	sink.Subscribe(ingress)
	for i := 0; i < 10; i++ {
		ingress <- []byte("test" + fmt.Sprint(i))
	}
}

func BenchmarkSyslog_Write(b *testing.B) {
	gen.RandomBytesLite(payload)
	for i := 0; i < b.N; i++ {
		_, err := benchWriter.Write(payload)
		if err != nil {
			continue
		}
	}
}

func BenchmarkSyslog_WriteParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		wr, _ := syslog.Dial("udp", "10.10.20.68:514", syslog.LOG_INFO, "test-bench")
		for pb.Next() {
			_, err := wr.Write(payload)
			if err != nil {
				continue
			}
		}
	})
}
