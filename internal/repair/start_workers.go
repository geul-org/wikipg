//ff:func feature=repair type=command control=iteration dimension=1
//ff:what startWorkers — starts c.Workers workers on the jobs and closes their outcome channel when all are done
package repair

import (
	"database/sql"
	"sync"
)

// startWorkers starts c.Workers workers (work) reading jobs and sending to a
// new outcome channel, which closes once every worker is done.
func startWorkers(db *sql.DB, c Config, jobs <-chan Span) <-chan outcome {
	out := make(chan outcome, c.Workers)
	var wg sync.WaitGroup
	for i := 0; i < c.Workers; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); work(db, c, jobs, out) }()
	}
	go func() { wg.Wait(); close(out) }()
	return out
}
