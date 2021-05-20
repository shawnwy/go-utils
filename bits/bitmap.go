package bits

import (
	"bytes"
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

type BitMap struct {
	m      []*bitset.BitSet
	width  uint
	height uint
}

func NewBitMap(width, height uint) *BitMap {
	m := make([]*bitset.BitSet, height)
	for h := uint(0); h < height; h++ {
		m[h] = bitset.New(uint(width))
	}
	return &BitMap{m, width, height}
}

// Set - Set <x, y> bit as 1.
// 	x must within [0, length)
//	y must within [0, width)
func (b *BitMap) Set(x, y uint) bool {
	if !b.withinRange(x, y) {
		return false
	}
	b.m[y].Set(x)
	return true
}

// Test - Check the bit <x, y> whether is 1 .
// 	x must within [0, length)
//	y must within [0, width)
func (b *BitMap) Test(x, y uint) bool {
	if !b.withinRange(x, y) {
		return false
	}
	return b.m[y].Test(x)
}

func (b *BitMap) withinRange(x, y uint) bool {
	if x < 0 || x >= b.width {
		return false
	}
	if y < 0 || y >= b.height {
		return false
	}
	return true
}

func (b *BitMap) DumpAsBits() string {
	buffer := bytes.NewBufferString("")
	for _, set := range b.m {
		fmt.Fprintln(buffer, set.DumpAsBits())
	}
	return buffer.String()
}
