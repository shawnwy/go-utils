package bits

import (
	"fmt"
	"testing"

	"github.com/bits-and-blooms/bitset"
)

func TestBitMap(t *testing.T) {
	m := NewBitMap(10, 5)
	fmt.Println(m)
	m.Set(5, 1)
	fmt.Println(m)

	s := bitset.New(8)
	fmt.Println(s)

	s.Set(1)
	fmt.Println(s)
	fmt.Println(s.Test(1), s.Len(), s.DumpAsBits())

	s.Set(0)
	fmt.Println(s)
	fmt.Println(s.Test(0), s.Len(), s.DumpAsBits())
}
