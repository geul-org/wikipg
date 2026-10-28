//ff:func feature=repair type=command control=iteration dimension=1
//ff:what work — one worker: checks each batch it is given (checkSpan) and sends the outcome, pausing after each
package repair

import (
	"database/sql"
	"time"
)

// work checks every batch from jobs (checkSpan, read only) and sends its
// outcome to out, then pauses c.Pause (the pace limit).
func work(db *sql.DB, c Config, jobs <-chan Span, out chan<- outcome) {
	for sp := range jobs {
		f, err := checkSpan(db, c, sp)
		f.Seq, f.Last = sp.Seq, sp.Last
		out <- outcome{Found: f, Err: err}
		time.Sleep(c.Pause)
	}
}
