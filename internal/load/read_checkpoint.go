//ff:func feature=load type=writer control=sequence
//ff:what ReadCheckpoint — the dump's committed-through line, its entity and error counts so far
package load

import (
	"database/sql"

	"github.com/lib/pq"
)

// ReadCheckpoint returns the line through which dump id is fully committed
// and the entity and error counts recorded with it (0 when none yet).
func ReadCheckpoint(db *sql.DB, schema string, id int16) (line, entities, errs int64, err error) {
	err = db.QueryRow("SELECT checkpoint_line, coalesce(entity_count, 0), coalesce(error_count, 0) FROM "+
		pq.QuoteIdentifier(schema)+".dump WHERE dump_id = $1", id).Scan(&line, &entities, &errs)
	return line, entities, errs, err
}
