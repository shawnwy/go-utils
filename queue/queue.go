package queue

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

type Queue struct {
	list *arraylist.List
}

func New() *Queue {
	return &Queue{list: arraylist.New()}
}

func (q *Queue) Size() int {
	return q.list.Size()
}

func (q *Queue) Empty() bool {
	return q.list.Empty()
}

func (q *Queue) Clear() {
	q.list.Clear()
}

func (q *Queue) withinRange(index int) bool {
	return index >= 0 && index < q.list.Size()
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.Add(value)
}

func (q *Queue) Dequeue() (value interface{}, ok bool) {
	value, ok = q.list.Get(0)
	q.list.Remove(0)
	return
}

func (q *Queue) Values() []interface{} {
	size := q.list.Size()
	elements := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		elements[i], _ = q.list.Get(i) // in reverse (LIFO)
	}
	return elements
}

// String returns a string representation of container
func (q *Queue) String() string {
	str := "ArrayQueue\n"
	values := []string{}
	for _, value := range q.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
