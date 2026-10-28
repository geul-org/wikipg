//ff:func feature=repair type=reader control=selection
//ff:what planSpans — the batches of a run after its checkpoint: line windows for a line range, line lists for an entity range
package repair

import "database/sql"

// planSpans returns the batches of c after line after (the checkpoint): for
// an entity range the lines of those entities (entityLines) in batches, for
// a line range the windows from max(c.From, after+1) to c.To or, when c.To is
// 0, the dump's last line (lastLine).
func planSpans(db *sql.DB, c Config, after int64) ([]Span, error) {
	switch {
	case c.Kind != "":
		lines, err := entityLines(db, c, after)
		return listSpans(lines, c.Batch), err
	case c.To != 0:
		return windowSpans(max(c.From, after+1), c.To, c.Batch), nil
	default:
		last, err := lastLine(db, c)
		return windowSpans(max(c.From, after+1), last, c.Batch), err
	}
}
