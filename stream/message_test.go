package stream

import (
	"fmt"
	"testing"
)

var (
	m          = make(RawMessage, 100)
	x IMessage = m
)

type _PB []byte

func (b _PB) UID() []byte {
	return b
}

func (b _PB) Bytes() []byte {
	return b
}

var _ IMessage = new(_PB)

func TestTypeAssertion(t *testing.T) {
	_, ok := x.(RawMessage)
	if !ok {
		t.FailNow()
	}

	pb, ok := x.(_PB)
	fmt.Println(pb, ok)
}

func TestIMessageByte(t *testing.T) {
	fmt.Printf("origin %p\n", m)
	fmt.Printf("iface %p\n", x)
	fmt.Printf("bytes %p\n", x.Bytes())
}
