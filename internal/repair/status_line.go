//ff:func feature=repair type=helper control=sequence
//ff:what statusLine — one progress line: batches and line applied with the share done, entities, MB, hash and occurrence verdicts, replacements, occurrence orders written, statement_reference differences, rates, ETA
package repair

import (
	"fmt"
	"time"
)

// statusLine describes st after elapsed: batches applied of planned (and the
// share done) and the last line, entities and raw MB, hash verdicts,
// replaced and skipped hashes, occurrence verdicts with the orders found and
// written (order build), statement_reference rows differing, entities and MB
// per second and the time left at that pace.
func statusLine(st Stats, elapsed time.Duration) string {
	sec := max(elapsed.Seconds(), 0.001)
	var eta time.Duration
	var done float64
	if st.Spans > 0 {
		eta = time.Duration(float64(elapsed) * float64(st.Planned-st.Spans) / float64(st.Spans)).Round(time.Second)
		done = 100 * float64(st.Spans) / float64(max(st.Planned, 1))
	}
	return fmt.Sprintf("batch %d/%d (%.2f%%) line %d: %d entities %.1f MB; hashes same %d order %d content %d missing %d; replaced %d skipped %d; "+
		"occurrences same %d order %d content %d missing %d, orders written %d; statement_reference diffs %d; %.0f entities/s %.2f MB/s; left %s",
		st.Spans, st.Planned, done, st.Last, st.Entities, float64(st.Bytes)/1e6, st.Hashes.Same, st.Hashes.Order, st.Hashes.Content, st.Hashes.Missing,
		st.Replaced, st.Skipped, st.Occurrences.Same, st.Occurrences.Order, st.Occurrences.Content, st.Occurrences.Missing, st.Written,
		st.RefDiffs, float64(st.Entities)/sec, float64(st.Bytes)/1e6/sec, eta)
}
