# go-utils
utilities for golang development

## JobQ
A queue for time job execution. It is also a retry queue.

### Example

```go
    sec := time.NewTicker(time.Second)
    t := time.NewTicker(10*time.Second)
    q := NewJobQ()
    for {
        select {
        case tick := <-sec.C:
            log.Printf("[%v] - Intel<%s>\n", tick, q.Intel())

        case job := <- q.JobChan:
            log.Println("An Job is about to execute", job)

        case <-t.C:
            jid := rand.Int()
			deltaSec := rand.Int63n(4) + 1
			ttl := deltaSec * int64(time.Second)
			ddl := time.Now().Add(time.Duration(deltaSec) * time.Second).UnixNano()

            log.Printf("[%v] - Add Job<%d, %d>\n", ts, jid, ttl)
            // add a job with deadline <ddl>
            q.AddJob(jid, ddl)
            /* or you could retry a job after a ttl
            q.Retry(jid, ttl)
            */
        }
    }
```