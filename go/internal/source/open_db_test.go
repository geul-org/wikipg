//ff:func feature=source type=loader control=sequence
//ff:what openDB — test fixture: the geuldev database from the PG* environment; skips the test when it cannot connect
package source

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

// openDB opens geuldev through lib/pq's PG* environment variables (PGHOST,
// PGUSER, PGPASSWORD, PGDATABASE, PGSSLMODE) and skips the test when the
// database is not reachable.
func openDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("postgres", "")
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		t.Skipf("geuldev not reachable: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}
