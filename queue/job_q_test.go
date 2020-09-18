package queue

import (
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
			// q.Retry(ddl, ttl)
			q.AddJob(ddl, ddl)

		case s := <-sec.C:
			log.Printf("[%v] - Intel<%s>\n", s, q.Intel())

		case job := <-q.JobChan:
			log.Println("-------------->", time.Unix(0, job.(int64)))

		case <-done:
			t.Stop()
		}
	}
}
