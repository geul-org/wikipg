//ff:func feature=repair type=writer control=sequence
//ff:what applyOrder — one batch of an order build: a dry run only counts, a writing run stores its occurrence orders and the checkpoint in one transaction
package repair

import "database/sql"

// applyOrder applies batch f of an order build, whose earlier batches are all
// applied: a dry run only counts it (Stats.addOrder); a writing run stores
// f.Orders in reference_snak_order (insertOrders) and moves the checkpoint to
// f.Last in one transaction, so a stopped run resumes exactly. The batch is
// all in or not at all.
func applyOrder(db *sql.DB, c Config, f Found, st *Stats) error {
	if c.DryRun {
		st.addOrder(f, false)
		return nil
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := insertOrders(tx, c, f.Orders); err != nil {
		return err
	}
	if err := saveCheckpoint(tx, c, f.Last); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	st.addOrder(f, true)
	return nil
}
