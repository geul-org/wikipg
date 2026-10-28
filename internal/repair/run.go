//ff:func feature=repair type=command control=sequence
//ff:what Run — one repair run of reference_snak over a range: lock and work tables when writing, plan batches after the checkpoint, check in parallel, apply in order
package repair

import (
	"context"
	"database/sql"
)

// Run repairs (or with c.DryRun only compares) reference_snak over the range
// of c. A writing run takes the advisory lock, creates the work tables
// (sql/repair_reference_snak.sql) and resumes after its checkpoint; a dry
// run writes nothing and starts at the range's beginning. Batches
// (planSpans) are read and checked by c.Workers workers in parallel
// (checkSpan, read only) and applied in line order (applyInOrder), so each
// hash is handled at its first occurrence in the range. ctx stops the run
// after the batches already read are applied; the checkpoint stays exact.
// It returns the totals and the first error, or ctx's error when stopped.
func Run(ctx context.Context, db *sql.DB, c Config) (Stats, error) {
	var after int64
	if !c.DryRun {
		conn, err := takeLock(ctx, db, c)
		if err != nil {
			return Stats{}, err
		}
		defer releaseLock(conn, c)
		if err := ensureTables(db, c); err != nil {
			return Stats{}, err
		}
		if after, err = readCheckpoint(db, c); err != nil {
			return Stats{}, err
		}
	}
	spans, err := planSpans(db, c, after)
	if err != nil {
		return Stats{}, err
	}
	st := Stats{Planned: int64(len(spans)), Last: after}
	inner, cancel := context.WithCancel(ctx)
	defer cancel()
	tokens, jobs := make(chan struct{}, 2*c.Workers), make(chan Span)
	go feed(inner, spans, tokens, jobs)
	if err := applyInOrder(db, c, startWorkers(db, c, jobs), tokens, cancel, &st); err != nil {
		return st, err
	}
	return st, ctx.Err()
}
