//ff:func feature=repair type=reader control=sequence
//ff:what snakList — a hash's stored reference_snak rows as one line: property=value/value_ordinal in ordinal order
package repair

import (
	"database/sql"
	"testing"
)

// snakList returns the stored rows of hash as "P=v/k …" in ordinal order, v
// the item ID, time, URL or snak type, k the value ordinal ("" for none).
func snakList(t *testing.T, db *sql.DB, c Config, hash string) string {
	t.Helper()
	var s sql.NullString
	err := db.QueryRow("SELECT string_agg(property || '=' || coalesce(value->>'id', value->>'time', value#>>'{}', snaktype::text) || '/' || value_ordinal, ' ' ORDER BY ordinal) FROM "+
		c.Schema+".reference_snak WHERE hash = $1", hash).Scan(&s)
	if err != nil {
		t.Fatal(err)
	}
	return s.String
}
