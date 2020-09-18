package queue

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type JobQ struct {
	*binaryheap.Heap
	*time.Ticker
	ddl     int64 // deadline nanoseconds
	quit    chan bool
	jobs    chan Element
	JobChan chan interface{}
	sync.RWMutex
}

type Element struct {
	val interface{}
	t   int64 // execution time in nanoseconds
}

func byExecTime(a, b interface{}) int {
	e1 := a.(*Element)
	e2 := b.(*Element)
	switch {
	case e1.t < e2.t:
		return -1
	case e1.t > e2.t:
		return 1
	default:
		return 0
	}
}

func NewJobQ() *JobQ {
	rq := &JobQ{
		Heap:    binaryheap.NewWith(byExecTime),
		Ticker:  time.NewTicker(24 * time.Hour),
		ddl:     time.Now().Add(24 * time.Hour).UnixNano(),
		JobChan: make(chan interface{}, 100000),
		quit:    make(chan bool, 1),
		jobs:    make(chan Element, 1),
		RWMutex: sync.RWMutex{},
	}
	rq.run()
	return rq
}

func (rq *JobQ) Intel() string {
	rq.RLock()
	defer rq.RUnlock()
	return fmt.Sprintf("ddl: %v(%d secs) queue<#%d>", rq.ddl/int64(time.Second), (time.Now().UnixNano()-rq.ddl)/int64(time.Second), rq.Size())
}

// Retry - Add a Job<e> with ttl(time-to-tile), duration for next execution.
// ttl is represented as nanoseconds
func (rq *JobQ) Retry(e interface{}, ttl int64) {
	ddl := time.Now().Add(time.Duration(ttl)).UnixNano()
	rq.jobs <- Element{e, ddl}
}

func (rq *JobQ) AddJob(e interface{}, ddl int64) {
	rq.jobs <- Element{e, ddl}
}

func (rq *JobQ) accelerate(ddl int64) {
	delta := math.Max(1.0, float64(ddl-time.Now().UnixNano()))
	rq.ddl = ddl
	rq.Ticker.Stop()
	rq.Ticker = time.NewTicker(time.Duration(delta))
	log.Printf("rerunning with %p ... next tick%v", rq.Ticker.C, delta/float64(time.Second))
}

func (rq *JobQ) execJobs() {
	for !rq.Empty() {
		v, _ := rq.Peek()
		e := v.(*Element)
		if e.t > time.Now().UnixNano() {
			return
		}
		rq.Pop()
		rq.JobChan <- e.val
	}
}

func (rq *JobQ) run() {
	go func() {
		for {
			select {
			case e := <-rq.jobs:
				// Exec Job<e> @ <ddl> nanoseconds
				var currDDL int64 = math.MaxInt64
				if v, ok := rq.Peek(); ok {
					top := v.(*Element)
					currDDL = top.t
				}
				rq.Push(&Element{e.val, e.t})
				fmt.Println(e.t, currDDL)
				if e.t < currDDL {
					rq.accelerate(e.t)
				}

			case t := <-rq.C:
				fmt.Println("ding @", t)
				rq.execJobs()
				fmt.Println("sent @", t)

			case <-rq.quit:
				rq.Stop()
				fmt.Println("terminating ...")
				return
			}
		}
	}()
}
