//ff:func feature=repair type=writer control=iteration dimension=1
//ff:what TestSaveCheckpoint — inserts, then moves the run's line; a missing table fails
package repair

import "testing"

func TestSaveCheckpoint(t *testing.T) {
	db, c := loadFixture(t)
	ensureTables(db, c)
	for _, line := range []int64{3, 5} {
		tx, _ := db.Begin()
		if err := saveCheckpoint(tx, c, line); err != nil {
			t.Fatal(err)
		}
		tx.Commit()
	}
	if n, _ := readCheckpoint(db, c); n != 5 {
		t.Fatalf("line %d", n)
	}
	c.Schema += "_none"
	tx, _ := db.Begin()
	defer tx.Rollback()
	if err := saveCheckpoint(tx, c, 1); err == nil {
		t.Fatal("missing table: want error")
	}
}
