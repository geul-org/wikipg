//ff:func feature=repair type=writer control=sequence
//ff:what ensureTables — creates the repair work tables (sql/repair_reference_snak.sql) in the schema if missing
package repair

import (
	"database/sql"

	"github.com/lib/pq"

	schema "github.com/geul-org/wikipg/sql"
)

// ensureTables runs sql/repair_reference_snak.sql (CREATE TABLE IF NOT
// EXISTS) as one implicit transaction with c.Schema as the search_path.
func ensureTables(db *sql.DB, c Config) error {
	_, err := db.Exec("SET LOCAL search_path = " + pq.QuoteIdentifier(c.Schema) + ";\n" + schema.RepairReferenceSnak)
	return err
}
