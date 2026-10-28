//ff:func feature=repair type=helper control=sequence
//ff:what statusLine — one progress line: batches and line applied, entities, MB, hash verdicts, replacements, statement_reference differences, rates, ETA
package repair

import (
	"fmt"
	"time"
)

// statusLine describes st after elapsed: batches applied of planned and the
// last line, entities and raw MB, hash verdicts, replaced and skipped hashes,
// statement_reference rows differing, entities and MB per second and the
// time left at that pace.
func statusLine(st Stats, elapsed time.Duration) string {
	sec := max(elapsed.Seconds(), 0.001)
	var eta time.Duration
	if st.Spans > 0 {
		eta = time.Duration(float64(elapsed) * float64(st.Planned-st.Spans) / float64(st.Spans)).Round(time.Second)
	}
	return fmt.Sprintf("batch %d/%d line %d: %d entities %.1f MB; hashes same %d order %d content %d missing %d; replaced %d skipped %d; statement_reference diffs %d; %.0f entities/s %.2f MB/s; left %s",
		st.Spans, st.Planned, st.Last, st.Entities, float64(st.Bytes)/1e6, st.Hashes.Same, st.Hashes.Order, st.Hashes.Content, st.Hashes.Missing,
		st.Replaced, st.Skipped, st.RefDiffs, float64(st.Entities)/sec, float64(st.Bytes)/1e6/sec, eta)
}
