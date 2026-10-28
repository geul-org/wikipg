//ff:func feature=repair type=reader control=iteration dimension=1
//ff:what readSpan — reads a batch's lines from entity_raw in line order and parses each as it streams (occurrencesOf)
package repair

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// readSpan reads the entity_raw lines of sp (a window: one range of the
// index entity_raw_dump_line; or a list of lines: one index lookup per line,
// LATERAL with OFFSET 0 as in fetchStored) in line order and parses each as it arrives (occurrencesOf), keeping only the
// occurrences and statement_reference rows, so memory holds one raw line at
// a time. A line that does not parse is kept as an error, not a failure.
func readSpan(db *sql.DB, c Config, sp Span) (spanData, error) {
	t := pq.QuoteIdentifier(c.Schema) + ".entity_raw"
	q, args := "SELECT line_no, raw FROM "+t+" WHERE dump_id = $1 AND line_no BETWEEN $2 AND $3 ORDER BY line_no", []any{c.DumpID, sp.From, sp.To}
	if sp.Lines != nil {
		q = "SELECT r.line_no, r.raw FROM unnest($2::bigint[]) AS l(line) CROSS JOIN LATERAL (SELECT line_no, raw FROM " + t +
			" WHERE dump_id = $1 AND line_no = l.line OFFSET 0) r ORDER BY r.line_no"
		args = []any{c.DumpID, pq.Array(sp.Lines)}
	}
	rows, err := db.Query(q, args...)
	if err != nil {
		return spanData{}, err
	}
	defer rows.Close()
	d := spanData{IDs: map[int64]string{}}
	for rows.Next() {
		var line int64
		var raw string
		if err := rows.Scan(&line, &raw); err != nil {
			return spanData{}, err
		}
		d.Lines, d.Bytes = append(d.Lines, line), d.Bytes+int64(len(raw))
		id, occs, refs, err := occurrencesOf(c.DumpID, line, raw)
		d.IDs[line] = id
		if err != nil {
			d.Errors = append(d.Errors, fmt.Sprintf("line %d %s: %v", line, id, err))
			continue
		}
		d.Occs, d.Refs = append(d.Occs, occs...), append(d.Refs, refs...)
	}
	return d, rows.Err()
}
