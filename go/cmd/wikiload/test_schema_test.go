//ff:func feature=cli type=command control=sequence
//ff:what testSchema — a throwaway copy of the wikidata schema in geuldev, dropped after the test; skips when geuldev is not reachable
package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// testSchema opens geuldev through lib/pq's PG* environment variables, runs
// sections 1–3 of wiki/wikidata.sql under a fresh schema name (wikidata is
// replaced) and drops that schema when the test ends. It skips the test when
// the database is not reachable.
func testSchema(t *testing.T) (*sql.DB, string) {
	t.Helper()
	db, err := sql.Open("postgres", "")
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		t.Skipf("geuldev not reachable: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	b, err := os.ReadFile("../../../wikidata.sql")
	if err != nil {
		t.Fatal(err)
	}
	name := fmt.Sprintf("wikidata_test_%d", time.Now().UnixNano())
	ddl := strings.ReplaceAll(string(b), "CREATE SCHEMA IF NOT EXISTS wikidata;\nSET search_path = wikidata;", "CREATE SCHEMA "+name+";\nSET search_path = "+name+";")
	if _, err := db.Exec(ddl); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Exec("DROP SCHEMA " + name + " CASCADE") })
	return db, name
}
