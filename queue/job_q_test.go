package queue

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestJobQAdd(test *testing.T) {
	done := make(chan os.Signal)
	t := time.NewTicker(1 * time.Second)
	sec := time.NewTicker(time.Second)
	q := NewJobQ()
	for {
		select {
		case ts := <-t.C:
			jid := rand.Int()
			deltaSec := rand.Int63n(4) + 1
			ttl := deltaSec * int64(time.Second)
			ddl := time.Now().Add(time.Duration(deltaSec) * time.Second).UnixNano()

			log.Printf("[%v] - Add Job<%d, %d>\n", ts, jid, ttl)
			q.AddJob(ddl, ddl)
			// jid := time.Now().Add(1).UnixNano()
			// log.Printf("[%v] - Add Job<%d, %d>\n", ts, jid, 1)
			// q.Retry(jid, 1)

		case s := <-sec.C:
			log.Printf("[%v] - Intel<%s>\n", s, q.Intel())

		case job := <-q.JobChan:
			log.Println("-------------->", job)
			log.Println("-------------->", time.Unix(0, job.(int64)))

		case <-done:
			t.Stop()
		}
	}
}

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func NewList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(val int) *LinkedList {
	if l.head == nil {
		l.head = &Node{val: val}
		l.tail = l.head
		l.len = 1
		return l
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
	l.len += 1
	return l
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Reverse() {
	if l.len <= 1 {
		return
	}
	dummy := &Node{next: l.head}
	l.tail = l.head
	for curr := dummy.next; curr.next != nil; {
		next := curr.next
		curr.next = next.next
		next.next = dummy.next
		dummy.next = next
	}
	l.head = dummy.next
}

func (l *LinkedList) String() {
	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Printf("-> %d ", curr.val)
	}
	fmt.Println("")
}

func TestReverse(t *testing.T) {
	l := NewList()
	l.String()
	l.Append(1).Append(2).Append(3).Append(4)
	l.String()
	l.Reverse()
	l.String()
}
