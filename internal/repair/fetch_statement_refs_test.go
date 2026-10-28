//ff:func feature=repair type=reader control=sequence
//ff:what TestFetchStatementRefs — the rows of the given lines only, in key order; a NULL hash and a missing table fail
package repair

import "testing"

func TestFetchStatementRefs(t *testing.T) {
	db, c := loadFixture(t)
	got, err := fetchStatementRefs(db, c, []int64{3, 2})
	if err != nil || len(got) != 5 || got[0] != (stmtRef{Statement: 1<<55 | 2<<20, Ordinal: 0, Hash: "hx"}) || got[4].Hash != "hx" {
		t.Fatalf("%v %v", got, err)
	}
	db.Exec("ALTER TABLE " + c.Schema + ".statement_reference ALTER hash DROP NOT NULL")
	db.Exec("INSERT INTO "+c.Schema+".statement_reference VALUES ($1, 0, NULL)", int64(1<<55|9<<20))
	if _, err := fetchStatementRefs(db, c, []int64{9}); err == nil {
		t.Fatal("NULL hash: want error")
	}
	if _, err := fetchStatementRefs(db, Config{Schema: c.Schema + "_none"}, []int64{2}); err == nil {
		t.Fatal("missing table: want error")
	}
}
