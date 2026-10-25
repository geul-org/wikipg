//ff:func feature=source type=loader control=iteration dimension=1
//ff:what loadTexts — the language-tagged strings a query returns as (language, text) rows
package source

import "database/sql"

// loadTexts runs query (which selects language and text for $1 = id among $2
// = langs) and returns the rows in order.
func loadTexts(db *sql.DB, query, id string, langs any) ([]Text, error) {
	rows, err := db.Query(query, id, langs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Text
	for rows.Next() {
		var t Text
		if err := rows.Scan(&t.Lang, &t.Value); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}
