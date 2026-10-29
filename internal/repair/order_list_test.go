//ff:func feature=repair type=reader control=sequence
//ff:what orderList — the stored reference_snak_order rows as one line: statement/ordinal=perm in key order
package repair

import (
	"database/sql"
	"testing"
)

// orderList returns the rows of reference_snak_order as "s/o={…} …" in key
// order, s the statement key's line and position (line.i), "" for none.
func orderList(t *testing.T, db *sql.DB, c Config) string {
	t.Helper()
	var s sql.NullString
	err := db.QueryRow("SELECT string_agg(((statement_id >> 20) & 34359738367) || '.' || (statement_id & 1048575) || '/' || ordinal || '=' || perm::text, ' ' ORDER BY statement_id, ordinal) FROM " +
		c.Schema + ".reference_snak_order").Scan(&s)
	if err != nil {
		t.Fatal(err)
	}
	return s.String
}
