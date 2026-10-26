//ff:func feature=load type=writer control=iteration dimension=1
//ff:what copyTable — COPYs rows into one table inside a transaction; nothing for no rows
package load

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// copyTable COPYs rows into schema.t within tx (lib/pq CopyInSchema). No rows
// is no statement.
func copyTable(tx *sql.Tx, schema string, t table, rows [][]any) error {
	if len(rows) == 0 {
		return nil
	}
	st, err := tx.Prepare(pq.CopyInSchema(schema, t.Name, t.Columns...))
	if err != nil {
		return err
	}
	for _, row := range rows {
		if _, err := st.Exec(row...); err != nil {
			st.Close()
			return fmt.Errorf("%s: %w", t.Name, err)
		}
	}
	if _, err := st.Exec(); err != nil {
		st.Close()
		return fmt.Errorf("%s: %w", t.Name, err)
	}
	return st.Close()
}
