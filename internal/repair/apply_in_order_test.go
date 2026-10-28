//ff:func feature=repair type=command control=sequence
//ff:what TestApplyInOrder — out-of-order outcomes apply in order with status lines; a failed batch or apply stops the run and later outcomes are only drained
package repair

import (
	"bytes"
	"errors"
	"testing"
	"time"
)

func TestApplyInOrder(t *testing.T) {
	db, c := loadFixture(t)
	var log bytes.Buffer
	c.Log, c.Status = &log, time.Nanosecond
	run := func(c Config, outs ...outcome) (Stats, error, bool) {
		results, tokens := make(chan outcome, len(outs)), make(chan struct{}, len(outs))
		for _, o := range outs {
			results <- o
			tokens <- struct{}{}
		}
		close(results)
		stopped := false
		var st Stats
		err := applyInOrder(db, c, results, tokens, func() { stopped = true }, &st)
		return st, err, stopped
	}
	st, err, stopped := run(c, outcome{Found: Found{Seq: 1, Last: 4}}, outcome{Found: Found{Seq: 0, Last: 2}})
	if err != nil || stopped || st.Spans != 2 || st.Last != 4 || log.Len() == 0 {
		t.Fatalf("in order: %+v %v %v", st, err, stopped)
	}
	st, err, stopped = run(c, outcome{Found: Found{Seq: 0}, Err: errors.New("read")}, outcome{Found: Found{Seq: 1}}, outcome{Found: Found{Seq: 2}, Err: errors.New("again")})
	if err == nil || !stopped || st.Spans != 0 {
		t.Fatalf("failed batch: %+v %v %v", st, err, stopped)
	}
	c.DryRun = false
	if _, err, stopped = run(c, outcome{Found: Found{Seq: 0}}); err == nil || !stopped {
		t.Fatalf("failed apply: %v %v", err, stopped)
	}
}
