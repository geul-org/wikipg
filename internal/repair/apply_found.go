//ff:func feature=repair type=writer control=sequence
//ff:what applyFound — applies one checked batch in line order: a dry run only counts, a writing run repairs (applyWrite)
package repair

import (
	"database/sql"

	"github.com/geul-org/wikipg/internal/load"
)

// applyFound applies batch f, whose earlier batches are all applied: the
// hashes new to this run (freshHashes) are counted in a dry run, and in a
// writing run repaired in one transaction with the checkpoint (applyWrite).
func applyFound(db *sql.DB, c Config, f Found, refs *load.RefSet, st *Stats) error {
	fresh := freshHashes(f, refs)
	if c.DryRun {
		st.add(f, fresh, 0, false)
		return nil
	}
	processed, skipped, err := applyWrite(db, c, f.Last, fresh)
	if err != nil {
		return err
	}
	st.add(f, processed, skipped, true)
	return nil
}
