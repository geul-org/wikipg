//ff:func feature=repair type=reader control=sequence
//ff:what TestPlanSpans — an entity range gives line lists, a closed line range windows after the checkpoint, an open one ends at the last line; a failed read fails
package repair

import "testing"

func TestPlanSpans(t *testing.T) {
	db, c := loadFixture(t)
	if got, err := planSpans(db, c, 0); err != nil || len(got) != 3 || got[0].From != 1 || got[2].To != 5 {
		t.Fatalf("open: %+v %v", got, err)
	}
	c.From, c.To = 2, 4
	if got, err := planSpans(db, c, 2); err != nil || len(got) != 1 || got[0].From != 3 || got[0].To != 4 {
		t.Fatalf("closed: %+v %v", got, err)
	}
	c.Kind, c.Lo, c.Hi = "Q", 1, 3
	if got, err := planSpans(db, c, 0); err != nil || len(got) != 2 || got[1].Lines[0] != 4 {
		t.Fatalf("entities: %+v %v", got, err)
	}
	c.Kind, c.To, c.Schema = "", 0, c.Schema+"_none"
	if _, err := planSpans(db, c, 0); err == nil {
		t.Fatal("missing table: want error")
	}
}
