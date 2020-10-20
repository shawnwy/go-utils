package queue

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

type Q struct {
	list *arraylist.List
}

func New() *Q {
	return &Q{list: arraylist.New()}
}

func (q *Q) Size() int {
	return q.list.Size()
}

func (q *Q) Empty() bool {
	return q.list.Empty()
}

func (q *Q) Clear() {
	q.list.Clear()
}

func (q *Q) withinRange(index int) bool {
	return index >= 0 && index < q.list.Size()
}

func (q *Q) Enque(value interface{}) {
	q.list.Add(value)
}

func (q *Q) Deque() (value interface{}, ok bool) {
	value, ok = q.list.Get(0)
	q.list.Remove(0)
	return
}

func (q *Q) Values() []interface{} {
	size := q.list.Size()
	elements := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		elements[i], _ = q.list.Get(i) // in reverse (LIFO)
	}
	return elements
}

// String returns a string representation of container
func (q *Q) String() string {
	str := "ArrayQ\n"
	values := []string{}
	for _, value := range q.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
