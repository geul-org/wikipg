//ff:func feature=repair type=reader control=sequence
//ff:what TestLastLine — the dump's last raw line; 0 for a dump without lines; a missing table fails
package repair

import "testing"

func TestLastLine(t *testing.T) {
	db, c := loadFixture(t)
	if n, err := lastLine(db, c); err != nil || n != 5 {
		t.Fatalf("%d %v", n, err)
	}
	c.DumpID = 2
	if n, err := lastLine(db, c); err != nil || n != 0 {
		t.Fatalf("empty: %d %v", n, err)
	}
	c.Schema += "_none"
	if _, err := lastLine(db, c); err == nil {
		t.Fatal("missing table: want error")
	}
}
