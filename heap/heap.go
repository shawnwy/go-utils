package heap

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

// Heap - basic data structure as min-heap
// with customized comparator and validator
type Heap struct {
	list arraylist.List
	utils.Comparator
	Validator
}

// Validator - It is element validator for identifing
// whether two element is same. Usually, the Validator
// is used for helping Find(), Update and Remove() implement
// their functionality.
type Validator func(a, b interface{}) bool

func NewHeap(c utils.Comparator, v Validator) *Heap {
	return &Heap{
		list:       arraylist.List{},
		Comparator: c,
		Validator:  v,
	}
}

// String - return Heap Contents in string
func (h *Heap) String() string {
	str := "BinaryHeap <"
	values := []string{}
	for _, value := range h.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ") + ">"
	return str
}

// Size - return the size of the heap
func (h *Heap) Size() int {
	return h.list.Size()
}

// Empty - return whether heap is empty
func (h *Heap) Empty() bool {
	return h.list.Size() == 0
}

// Push - pushes the element 'e' onto the heap
// The complexity is O(log n) where n = h.Len()
func (h *Heap) Push(e interface{}) {
	h.list.Add(e)
	h.siftUp()
}

// Pop removes and returns the minimum element (according to Comparator) from heap
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap) Pop() (interface{}, bool) {
	val, ok := h.list.Get(0)
	if !ok {
		return nil, false
	}
	lastIdx := h.list.Size() - 1
	h.list.Swap(0, lastIdx)
	h.list.Remove(lastIdx)
	h.siftDown()
	return val, true
}

// Peek - return the minimum element (according to Comparator) from heap
// The complexity is O(1)
func (h *Heap) Peek() (interface{}, bool) {
	return h.list.Get(0)
}

// Remove - removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Remove(idx int) {
	if !h.withinRange(idx) {
		return
	}
	lastIdx := h.list.Size() - 1
	if lastIdx == idx {
		h.list.Remove(lastIdx)
		return
	}
	h.list.Swap(idx, lastIdx)
	h.list.Remove(lastIdx)
	h.siftDownIndex(idx)
}

// Update - update an existed element by index
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Update(e interface{}, idx int) bool {
	if idx < 0 {
		return false
	}
	h.list.Set(idx, e)
	h.list.Swap(idx, h.list.Size()-1)
	if !h.siftDownIndex(idx) {
		h.siftUpIndex(idx)
	}
	return true
}

// Find - find the certain element by cutomized validation function
// return (index, value) if the element existed in the heap
// return (-1, nil) if doesn't existed
// the complexity is O(n) where h = h.Len()
func (h *Heap) Find(e interface{}) (int, interface{}) {
	return h.list.Find(func(index int, value interface{}) bool {
		return h.Validator(value, e)
	})
}

func (h *Heap) withinRange(idx int) bool {
	return idx >= 0 && idx < h.list.Size()
}

// Perform the "sift up" operation. This is to place a newly inserted
// element (i.e. last element in the list) in its correct place
func (h *Heap) siftUp() {
	idx := h.Size() - 1
	h.siftUpIndex(idx)
}

func (h *Heap) siftUpIndex(idx int) {
	for parentIdx := (idx - 1) >> 1; idx > 0; parentIdx = (idx - 1) >> 1 {
		idxVal, _ := h.list.Get(idx)
		parentVal, _ := h.list.Get(parentIdx)
		if h.Comparator(parentVal, idxVal) <= 0 {
			break
		}
		h.list.Swap(idx, parentIdx)
		idx = parentIdx
	}
}

// Perform the "sift down" operation. This is to place a newly inserted
// element in its correct place
func (h *Heap) siftDown() {
	h.siftDownIndex(0)
}

func (h *Heap) siftDownIndex(i0 int) bool {
	idx := i0
	size := h.Size()
	for leftIdx := idx<<1 + 1; leftIdx < size; leftIdx = idx<<1 + 1 {
		rightIdx := idx<<1 + 2
		smallerIdx := leftIdx
		leftVal, _ := h.list.Get(leftIdx)
		rightVal, _ := h.list.Get(rightIdx)
		if rightIdx < size && h.Comparator(leftVal, rightVal) > 0 {
			smallerIdx = rightIdx
		}
		idxVal, _ := h.list.Get(idx)
		smallerVal, _ := h.list.Get(smallerIdx)
		if h.Comparator(idxVal, smallerVal) <= 0 {
			break
		}
		h.list.Swap(idx, smallerIdx)
		idx = smallerIdx
	}
	return idx > i0
}
