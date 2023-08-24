package sinks

import (
	"fmt"
	"testing"
	"time"

	"github.com/shawnwy/go-utils/v5/gen"
)

type IMsg interface {
	UID() []byte
	Bytes() []byte
}

type _PB []byte

func (b _PB) UID() []byte {
	return b
}

func (b _PB) Bytes() []byte {
	return b
}

func TestChanWithIface(t *testing.T) {
	pb := make(_PB, 10)
	gen.RandomBytes(pb)
	ch := make(chan IMsg, 1)
	fmt.Printf("%s, %p\n", string(pb.Bytes()), pb)
	ch <- pb
	time.Sleep(time.Second)
	x := <-ch
	fmt.Printf("%s, %p\n", string(x.Bytes()), x)
}
