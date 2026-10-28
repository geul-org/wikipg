//ff:func feature=repair type=reader control=sequence
//ff:what TestReadCheckpoint — 0 before a run, the saved line after, per run name; a missing table fails
package repair

import "testing"

func TestReadCheckpoint(t *testing.T) {
	db, c := loadFixture(t)
	if err := ensureTables(db, c); err != nil {
		t.Fatal(err)
	}
	if n, err := readCheckpoint(db, c); err != nil || n != 0 {
		t.Fatalf("none: %d %v", n, err)
	}
	tx, _ := db.Begin()
	saveCheckpoint(tx, c, 4)
	tx.Commit()
	if n, err := readCheckpoint(db, c); err != nil || n != 4 {
		t.Fatalf("saved: %d %v", n, err)
	}
	other := c
	other.Kind, other.Lo, other.Hi = "Q", 1, 2
	if n, err := readCheckpoint(db, other); err != nil || n != 0 {
		t.Fatalf("other run: %d %v", n, err)
	}
	c.Schema += "_none"
	if _, err := readCheckpoint(db, c); err == nil {
		t.Fatal("missing table: want error")
	}
}
