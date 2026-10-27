//ff:func feature=cli type=command control=sequence
//ff:what fixtureDB — test fixture: a throwaway schema of sql/wikidata.sql holding testdata/fixture.sql, dropped after the test; skips when the database is not reachable
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

// fixtureDB opens the database of lib/pq's PG* environment variables, runs
// sections 1–3 of sql/wikidata.sql and testdata/fixture.sql under a fresh
// schema, drops it when the test ends and returns the schema name. It skips
// the test when the database is not reachable.
func fixtureDB(t *testing.T) (*sql.DB, string) {
	t.Helper()
	db, err := sql.Open("postgres", "")
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		t.Skipf("database not reachable: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	ddl, err := os.ReadFile("../../sql/wikidata.sql")
	if err != nil {
		t.Fatal(err)
	}
	fix, err := os.ReadFile("../../testdata/fixture.sql")
	if err != nil {
		t.Fatal(err)
	}
	name := fmt.Sprintf("wikiport_test_%d", time.Now().UnixNano())
	set := "CREATE SCHEMA " + name + ";\nSET search_path = " + name + ";"
	script := strings.Replace(string(ddl), "CREATE SCHEMA IF NOT EXISTS wikidata;\nSET search_path = wikidata;", set, 1) + "\nSET search_path = " + name + ";\n" + string(fix)
	if _, err := db.Exec(script); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Exec("DROP SCHEMA " + name + " CASCADE") })
	return db, name
}
