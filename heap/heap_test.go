package heap

import (
	"fmt"
	"testing"
)

type E struct {
	ID       int
	Priority float64
}

func TestHeap(t *testing.T) {
	h := NewHeap(
		func(a, b interface{}) int {
			x := a.(E)
			y := b.(E)
			switch {
			case x.Priority < y.Priority:
				return -1
			case x.Priority > y.Priority:
				return 1
			default:
				return 0
			}
		}, func(a, b interface{}) bool {
			x := a.(E)
			y := b.(E)
			return x.ID == y.ID
		})
	h.Push(E{1, 0.5})
	h.Push(E{2, 11})
	h.Push(E{3, 10})

	fmt.Println(h)
	idx, val := h.Find(E{3, -1})
	fmt.Println(idx, val)
	h.Remove(idx)
	fmt.Println(h)
	h.Push(E{4, 10})
	idx, val = h.Find(E{4, -1})
	fmt.Println(h)
	fmt.Println(idx, val)
	h.Update(E{4, 0.3}, idx)
	fmt.Println(h)
	// fmt.Println(h.Size())
	// val, exist := h.Pop()
	// fmt.Println(h.Size())

	for !h.Empty() {
		val, _ := h.Pop()
		fmt.Println(val)
		fmt.Println(val, h, h.Size())
	}

}
