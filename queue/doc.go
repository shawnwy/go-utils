// Package queue implements a thread-safe JobQ and ordinary Queue:
// 	JobQ is for job execution with ddl(deadline)
// 	which is based on [GoDs](https://github.com/emirpasic/gods) from emirpasic.

// 	JobQ is also a retry queue:
// 	- Add(job, ddl)
// 	- Retry(j, ttl)
// 	It use a dynamic ticker to re-adjust next tick.
package queue
