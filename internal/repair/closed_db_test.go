//ff:func feature=repair type=writer control=sequence
//ff:what closedDB — a lib/pq handle closed before use, for error paths
package repair

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

// closedDB returns a lib/pq handle closed before use: every call on it fails,
// which exercises the error paths.
func closedDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("postgres", "")
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	return db
}
