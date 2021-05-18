package bits

import "github.com/bits-and-blooms/bitset"

type BitMap struct {
	m      []*bitset.BitSet
	width  int
	length int
}

func NewBitMap(width, length int) *BitMap {
	m := make([]*bitset.BitSet, width)
	for w := 0; w < width; w++ {
		m[w] = bitset.New(uint(length))
	}
	return &BitMap{m, width, length}
}

// Set - Set <x, y> bit as 1.
// 	x must within [0, length)
//	y must within [0, width)
func (b *BitMap) Set(x, y int) bool {
	if !b.withInRange(x, y) {
		return false
	}
	b.m[y].Set(uint(x))
	return true
}

// Test - Check the bit <x, y> whether is 1 .
// 	x must within [0, length)
//	y must within [0, width)
func (b *BitMap) Test(x, y int) bool {
	if !b.withInRange(x, y) {
		return false
	}
	return b.m[y].Test(uint(x))
}

func (b *BitMap) withInRange(x, y int) bool {
	if x < 0 || x >= b.length {
		return false
	}
	if y < 0 || y >= b.width {
		return false
	}
	return true
}
