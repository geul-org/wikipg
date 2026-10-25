//ff:func feature=cli type=command control=sequence
//ff:what openDB — test fixture: the geuldev database from the PG* environment; skips the test when it cannot connect
package main

import (
	"database/sql"
	"testing"
)

// openDB opens geuldev through lib/pq's PG* environment variables and skips
// the test when the database is not reachable.
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
