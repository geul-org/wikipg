//ff:func feature=repair type=reader control=sequence
//ff:what readCheckpoint — the last line a run has applied (0 when it has not started)
package repair

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// readCheckpoint returns the line through which run c.Run() of the dump has
// been applied, 0 when there is no checkpoint yet.
func readCheckpoint(db *sql.DB, c Config) (int64, error) {
	var line int64
	err := db.QueryRow("SELECT line_no FROM "+pq.QuoteIdentifier(c.Schema)+".repair_reference_snak_checkpoint WHERE run = $1 AND dump_id = $2", c.Run(), c.DumpID).Scan(&line)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	return line, err
}
