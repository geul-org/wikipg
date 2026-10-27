//ff:func feature=source type=loader control=sequence
//ff:what entityHead — whether the head query finds the entity, and its datatype (properties only)
package source

import "database/sql"

// entityHead runs query (count and datatype for $1 = id) and reports whether
// the entity was loaded (a deleted or never created ID has no row) and its
// datatype: the Wikibase datatype of a property (url, external-id …), empty
// for an item.
func entityHead(db *sql.DB, query, id string) (string, bool, error) {
	var n int
	var datatype string
	if err := db.QueryRow(query, id).Scan(&n, &datatype); err != nil {
		return "", false, err
	}
	return datatype, n > 0, nil
}
