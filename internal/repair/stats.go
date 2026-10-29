//ff:type feature=repair type=model
//ff:what Stats — the totals of a repair or order build run: batches, lines, entities, bytes, hash and occurrence verdicts, replaced and skipped hashes, occurrence orders found and written, statement_reference differences, examples
package repair

import "time"

// Stats sums a run: the batches applied of all planned, the last line
// applied, entities and raw bytes read, the verdicts of each hash at its first
// occurrence in the run (Hashes), every occurrence compared with the rows read
// (Occurrences) and the entities with a content or an order difference,
// hashes replaced and hashes skipped as done in an earlier run, the
// occurrences in another order found and written (order build), the
// statement_reference rows compared and those differing with their entities,
// parse errors, up to 20 examples of differing hashes and the elapsed time.
type Stats struct {
	Spans           int64
	Planned         int64
	Last            int64
	Entities        int64
	Bytes           int64
	Hashes          Tally
	Occurrences     Tally
	ContentEntities int64
	OrderEntities   int64
	Replaced        int64
	Skipped         int64
	Orders          int64
	Written         int64
	StatementRefs   int64
	RefDiffs        int64
	RefDiffEntities []string
	Errors          []string
	Examples        []string
	Elapsed         time.Duration
}
