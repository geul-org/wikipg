//ff:func feature=repair type=writer control=sequence
//ff:what saveCheckpoint — records in a transaction that the run of the dump has applied through a line (key dump_id, run)
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// saveCheckpoint records in tx that run c.Run() of dump c.DumpID has applied
// every batch through line, in the checkpoint table of c's mode
// (Config.tables). The key is (dump_id, run): a range name is the same for
// every dump (lines:0-0), so another dump never moves this one's line.
func saveCheckpoint(tx *sql.Tx, c Config, line int64) error {
	_, table := c.tables()
	_, err := tx.Exec("INSERT INTO "+pq.QuoteIdentifier(c.Schema)+"."+table+" (dump_id, run, line_no) VALUES ($1, $2, $3) "+
		"ON CONFLICT (dump_id, run) DO UPDATE SET line_no = EXCLUDED.line_no, updated = now()", c.DumpID, c.Run(), line)
	return err
}
