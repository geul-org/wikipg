//ff:func feature=repair type=reader control=sequence
//ff:what readCheckpoint — the last line the run of the dump has applied (0 when it has not started)
package repair

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// readCheckpoint returns the line through which run c.Run() of dump c.DumpID
// has been applied, 0 when there is no checkpoint yet, from the checkpoint
// table of c's mode (Config.tables), keyed by (dump_id, run).
func readCheckpoint(db *sql.DB, c Config) (int64, error) {
	var line int64
	_, table := c.tables()
	err := db.QueryRow("SELECT line_no FROM "+pq.QuoteIdentifier(c.Schema)+"."+table+" WHERE dump_id = $1 AND run = $2", c.DumpID, c.Run()).Scan(&line)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	return line, err
}
