//ff:func feature=load type=writer control=sequence
//ff:what FinishDump — marks the dump's load finished now
package load

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// FinishDump sets load_finished of dump id to now: every line was read and
// committed.
func FinishDump(db *sql.DB, schema string, id int16) error {
	_, err := db.Exec("UPDATE "+pq.QuoteIdentifier(schema)+".dump SET load_finished = $2 WHERE dump_id = $1", id, time.Now())
	return err
}
