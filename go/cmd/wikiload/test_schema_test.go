//ff:func feature=cli type=command control=sequence
//ff:what testSchema — a throwaway schema of wiki/wikidata.sql, dropped after the test; skips when the database is not reachable
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

// testSchema opens the database of lib/pq's PG* environment variables, runs
// sections 1–3 of wiki/wikidata.sql under a fresh schema (the schema wikidata
// is replaced) and drops that schema when the test ends. It skips the test when
// the database is not reachable.
func testSchema(t *testing.T) (*sql.DB, string) {
	t.Helper()
	db, err := sql.Open("postgres", "")
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		t.Skipf("database not reachable: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	b, err := os.ReadFile("../../../wikidata.sql")
	if err != nil {
		t.Fatal(err)
	}
	name := fmt.Sprintf("wikidata_test_%d", time.Now().UnixNano())
	ddl := strings.Replace(string(b), "CREATE SCHEMA IF NOT EXISTS wikidata;\nSET search_path = wikidata;", "CREATE SCHEMA "+name+";\nSET search_path = "+name+";", 1)
	if _, err := db.Exec(ddl); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Exec("DROP SCHEMA " + name + " CASCADE") })
	return db, name
}
