package queue

import (
	"fmt"
	"testing"
)

func TestEnque(t *testing.T) {
	q := New()
	q.Enqueue(1)
	fmt.Println(q)
}

func TestDeque(t *testing.T) {
	q := New()
	q.Enqueue("a")
	q.Enqueue(1)
	fmt.Println(q)
	e, _ := q.Dequeue()
	fmt.Println(e)
}
