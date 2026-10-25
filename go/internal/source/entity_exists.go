//ff:func feature=source type=loader control=sequence
//ff:what entityExists — whether the existence query counts any row for the entity ID
package source

import "database/sql"

// entityExists runs query (a count for $1 = id) and reports whether the
// entity was loaded (a deleted or never created ID has no row).
func entityExists(db *sql.DB, query, id string) (bool, error) {
	var n int
	if err := db.QueryRow(query, id).Scan(&n); err != nil {
		return false, err
	}
	return n > 0, nil
}
