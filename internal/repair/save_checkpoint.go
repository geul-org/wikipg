//ff:func feature=repair type=writer control=sequence
//ff:what saveCheckpoint — records in a transaction that the run has applied through a line
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// saveCheckpoint records in tx that run c.Run() has applied every batch
// through line.
func saveCheckpoint(tx *sql.Tx, c Config, line int64) error {
	_, err := tx.Exec("INSERT INTO "+pq.QuoteIdentifier(c.Schema)+".repair_reference_snak_checkpoint (run, dump_id, line_no) VALUES ($1, $2, $3) "+
		"ON CONFLICT (run) DO UPDATE SET line_no = EXCLUDED.line_no, updated = now()", c.Run(), c.DumpID, line)
	return err
}
