//ff:func feature=repair type=command control=iteration dimension=1
//ff:what feed — hands the batches to the workers in order, at most as many ahead of the applier as there are slots, until done or stopped
package repair

import "context"

// feed sends spans to jobs in order, taking a slot of tokens for each (the
// applier frees it when the batch is applied), so at most cap(tokens)
// batches are read but not yet applied. It stops when ctx is done and
// closes jobs.
func feed(ctx context.Context, spans []Span, tokens chan<- struct{}, jobs chan<- Span) {
	defer close(jobs)
	for _, sp := range spans {
		select {
		case tokens <- struct{}{}:
		case <-ctx.Done():
			return
		}
		select {
		case jobs <- sp:
		case <-ctx.Done():
			return
		}
	}
}
