//ff:func feature=repair type=writer control=sequence
//ff:what TestReplaceSnaks — a hash's stored rows give way to its occurrence's rows; none is a no-op; a bad snak type and a missing table fail
package repair

import (
	"testing"
)

func TestReplaceSnaks(t *testing.T) {
	db, c := loadFixture(t)
	row := []any{"hx", int16(1), 0, "P854", 0, "value", "url", "string", `"https://example.org/n"`, nil, nil}
	tx, _ := db.Begin()
	if err := replaceSnaks(tx, c, []HashCheck{{Hash: "hx", Rows: [][]any{row}}, {Hash: "hz"}}); err != nil {
		t.Fatal(err)
	}
	if err := replaceSnaks(tx, c, nil); err != nil {
		t.Fatal(err)
	}
	tx.Commit()
	if got := snakList(t, db, c, "hx"); got != "P854=https://example.org/n/0" {
		t.Fatalf("hx: %q", got)
	}
	tx, _ = db.Begin()
	if err := replaceSnaks(tx, c, []HashCheck{{Hash: "hw", Rows: [][]any{{"hw", int16(1), 0, "P1", 0, "bogus", nil, nil, nil, nil, nil}}}}); err == nil {
		t.Fatal("bad snak type: want error")
	}
	tx.Rollback()
	bad := c
	bad.Schema += "_none"
	tx, _ = db.Begin()
	if err := replaceSnaks(tx, bad, []HashCheck{{Hash: "hw"}}); err == nil {
		t.Fatal("missing table: want error")
	}
	tx.Rollback()
}
