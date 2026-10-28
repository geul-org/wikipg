//ff:func feature=repair type=reader control=iteration dimension=1
//ff:what normalizeValues — each datavalue JSON text in the form PostgreSQL prints it as jsonb, so dump values compare with stored ones
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// normalizeValues maps each JSON text of values to its jsonb text form
// (v::jsonb::text: key order, spacing, number and escape forms as stored), in
// one read-only query keeping the input order.
func normalizeValues(db *sql.DB, values []string) (map[string]string, error) {
	var forms []string
	err := db.QueryRow("SELECT coalesce(array_agg(v::jsonb::text ORDER BY n), '{}') FROM unnest($1::text[]) WITH ORDINALITY AS t(v, n)", pq.Array(values)).Scan(pq.Array(&forms))
	if err != nil {
		return nil, err
	}
	out := make(map[string]string, len(values))
	for i, v := range values {
		out[v] = forms[i]
	}
	return out, nil
}
