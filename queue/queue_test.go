package queue

import (
	"fmt"
	"testing"
)

func TestEnque(t *testing.T) {
	q := New()
	q.Enque(1)
	fmt.Println(q)
}

func TestDeque(t *testing.T) {
	q := New()
	q.Enque("a")
	q.Enque(1)
	fmt.Println(q)
	e, _ := q.Deque()
	fmt.Println(e)
}
