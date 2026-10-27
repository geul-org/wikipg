//ff:func feature=load type=writer control=iteration dimension=1
//ff:what WriteRows — writes one batch in one transaction: a COPY per table that has rows
package load

import "database/sql"

// WriteRows writes the batch r into schema in one transaction, one COPY per
// table with rows in tables order, so a batch is either all in or not at all.
// The commit does not wait for the WAL flush (synchronous_commit off): a crash
// can lose the last moments of commits but never leaves a batch half in, and
// a checkpoint saved after them is later in the WAL, so it is lost with them.
func WriteRows(db *sql.DB, schema string, r *Rows) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec("SET LOCAL synchronous_commit = off"); err != nil {
		tx.Rollback()
		return err
	}
	for _, t := range tables() {
		if err := copyTable(tx, schema, t, r.byTable[t.Name]); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
