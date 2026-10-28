//ff:func feature=repair type=reader control=sequence
//ff:what lastLine — the last line of the dump in entity_raw (0 when none)
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// lastLine returns the largest entity_raw line of the dump (index
// entity_raw_dump_line), 0 when it has none.
func lastLine(db *sql.DB, c Config) (int64, error) {
	var n sql.NullInt64
	err := db.QueryRow("SELECT max(line_no) FROM "+pq.QuoteIdentifier(c.Schema)+".entity_raw WHERE dump_id = $1", c.DumpID).Scan(&n)
	return n.Int64, err
}
