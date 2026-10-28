//ff:func feature=repair type=writer control=sequence
//ff:what applyWrite — one batch's repair in one transaction: skip hashes done, replace differing rows, record hashes done and the checkpoint
package repair

import "database/sql"

// applyWrite applies the fresh checks of a batch ending at line last in one
// transaction: hashes already done at an earlier or the same occurrence are
// skipped (doneOf, filterDone), the rows of the rest that differ are replaced
// (replaceSnaks), the rest are recorded as done (markDone) and the checkpoint
// moves to last. It returns the checks processed and the number skipped. The
// batch is all in or not at all.
func applyWrite(db *sql.DB, c Config, last int64, fresh []HashCheck) ([]HashCheck, int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	done, err := doneOf(tx, c, hashesOf(fresh))
	if err != nil {
		return nil, 0, err
	}
	process, skipped := filterDone(fresh, done)
	if err := replaceSnaks(tx, c, differing(process)); err != nil {
		return nil, 0, err
	}
	if err := markDone(tx, c, process); err != nil {
		return nil, 0, err
	}
	if err := saveCheckpoint(tx, c, last); err != nil {
		return nil, 0, err
	}
	return process, skipped, tx.Commit()
}
