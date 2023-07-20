package flags

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const (
	CommKafka = "kafka"
	CommNATs  = "nats"
)

type BaseSetting struct {
	Pprof string `json:"http-pprof"`
}

func (b BaseSetting) IsPprofON() bool {
	return b.Pprof != ""
}

func (b BaseSetting) AttachPprof(h http.Handler) {
	if b.IsPprofON() {
		go func() {
			log.Println(http.ListenAndServe(b.Pprof, h))
		}()
	}
}

func DefineBaseFlags(b *BaseSetting, cmd string) {
	flag.StringVar(&b.Pprof, "http-pprof", "", fmt.Sprintf("Enable <%s> profiling. Starts  http server on specified port, exposing special /debug/pprof endpoint. Example: `:8989`", cmd))
}
