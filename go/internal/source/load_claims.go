//ff:func feature=source type=loader control=iteration dimension=1
//ff:what loadClaims — the triples a query returns as (id, property, type, value, rank) rows
package source

import "database/sql"

// loadClaims runs query (which selects triple ID, property, value type, value
// and rank for $1 = id) and returns the rows in order.
func loadClaims(db *sql.DB, query, id string) ([]Claim, error) {
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Claim
	for rows.Next() {
		var c Claim
		if err := rows.Scan(&c.ID, &c.Property, &c.Type, &c.Value, &c.Rank); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}
