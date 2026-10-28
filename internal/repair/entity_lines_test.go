//ff:func feature=repair type=reader control=sequence
//ff:what TestEntityLines — the lines of an entity range after a line, in line order; none is empty; a missing table fails
package repair

import (
	"slices"
	"testing"
)

func TestEntityLines(t *testing.T) {
	db, c := loadFixture(t)
	c.Kind, c.Lo, c.Hi = "Q", 2, 4
	if got, err := entityLines(db, c, 3); err != nil || !slices.Equal(got, []int64{4, 5}) {
		t.Fatalf("%v %v", got, err)
	}
	if got, err := entityLines(db, c, 5); err != nil || len(got) != 0 {
		t.Fatalf("none: %v %v", got, err)
	}
	c.Schema += "_none"
	if _, err := entityLines(db, c, 0); err == nil {
		t.Fatal("missing table: want error")
	}
}
