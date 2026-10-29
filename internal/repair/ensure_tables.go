//ff:func feature=repair type=writer control=sequence
//ff:what ensureTables — creates the work tables of the run's mode (sql/repair_reference_snak.sql or sql/reference_snak_order.sql) in the schema if missing
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// ensureTables runs the SQL of c's mode (Config.tables: CREATE TABLE IF NOT
// EXISTS) as one implicit transaction with c.Schema as the search_path.
func ensureTables(db *sql.DB, c Config) error {
	ddl, _ := c.tables()
	_, err := db.Exec("SET LOCAL search_path = " + pq.QuoteIdentifier(c.Schema) + ";\n" + ddl)
	return err
}
