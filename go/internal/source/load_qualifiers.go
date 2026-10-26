//ff:func feature=source type=loader control=iteration dimension=1
//ff:what loadQualifiers — the qualifiers of an entity's claims a query returns as (statement ID, property, datatype, value) rows
package source

import "database/sql"

// loadQualifiers runs query (which selects statement ID, property, value type, datatype and
// value of the qualifiers of $1 = id's claims) and returns the rows in order.
func loadQualifiers(db *sql.DB, query, id string) ([]Qualifier, error) {
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Qualifier
	for rows.Next() {
		var q Qualifier
		if err := rows.Scan(&q.StatementID, &q.Property, &q.Type, &q.Datatype, &q.Value); err != nil {
			return nil, err
		}
		out = append(out, q)
	}
	return out, rows.Err()
}
