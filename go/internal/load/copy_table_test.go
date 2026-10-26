//ff:func feature=load type=writer control=sequence
//ff:what TestCopyTable — rows COPY in; no rows is nothing; a short row, a bad or unconvertible value and a bad table fail
package load

import "testing"

func TestCopyTable(t *testing.T) {
	db, schema := testSchema(t)
	label := table{"label", []string{"entity_id", "dump_id", "language", "value"}}
	tx, _ := db.Begin()
	if err := copyTable(tx, schema, label, [][]any{{"Q1", 1, "en", "a"}, {"Q2", 1, "en", "b"}}); err != nil {
		t.Fatal(err)
	}
	if err := copyTable(tx, schema, label, nil); err != nil {
		t.Fatal(err)
	}
	var n int
	tx.QueryRow("SELECT count(*) FROM " + schema + ".label").Scan(&n)
	tx.Rollback()
	if n != 2 {
		t.Fatalf("rows %d", n)
	}
	tx, _ = db.Begin()
	if err := copyTable(tx, schema, label, [][]any{{"Q1", 1, "en"}}); err == nil {
		t.Fatal("short row: want error")
	}
	tx.Rollback()
	tx, _ = db.Begin()
	if err := copyTable(tx, schema, label, [][]any{{"Q1", "not a number", "en", "a"}}); err == nil {
		t.Fatal("bad value: want error")
	}
	tx.Rollback()
	tx, _ = db.Begin()
	if err := copyTable(tx, schema, label, [][]any{{"Q1", 1, "en", struct{}{}}}); err == nil {
		t.Fatal("unconvertible value: want error")
	}
	tx.Rollback()
	tx, _ = db.Begin()
	if err := copyTable(tx, schema, table{"nope", []string{"x"}}, [][]any{{1}}); err == nil {
		t.Fatal("missing table: want error")
	}
	tx.Rollback()
}
