//ff:func feature=repair type=reader control=sequence
//ff:what entityLines — the dump lines of the entities of one kind in a number range, in line order
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// entityLines returns the lines of the entities c.Kind c.Lo..c.Hi of the dump
// (entity, index entity_dump_kind_num) after line after, in line order, as
// one array: an entity range is meant for a part of the dump (Q1-Q10000:
// 9,656 lines); the whole dump goes by lines.
func entityLines(db *sql.DB, c Config, after int64) ([]int64, error) {
	var lines []int64
	err := db.QueryRow("SELECT coalesce(array_agg(line_no ORDER BY line_no), '{}') FROM "+pq.QuoteIdentifier(c.Schema)+".entity "+
		"WHERE dump_id = $1 AND kind = $2 AND num BETWEEN $3 AND $4 AND line_no > $5", c.DumpID, c.Kind, c.Lo, c.Hi, after).Scan(pq.Array(&lines))
	return lines, err
}
