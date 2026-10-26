//ff:func feature=load type=writer control=sequence
//ff:what SaveCheckpoint — records that every line through a number is committed, with the counts so far
package load

import (
	"database/sql"

	"github.com/lib/pq"
)

// SaveCheckpoint records on dump id that every line through line is committed
// and the entity and error counts so far.
func SaveCheckpoint(db *sql.DB, schema string, id int16, line, entities, errs int64) error {
	_, err := db.Exec("UPDATE "+pq.QuoteIdentifier(schema)+".dump SET checkpoint_line = $2, entity_count = $3, error_count = $4 WHERE dump_id = $1",
		id, line, entities, errs)
	return err
}
