//ff:func feature=repair type=helper control=sequence
//ff:what TestWindowSpans — windows cover the range in order with a short last one; an empty range gives none
package repair

import "testing"

func TestWindowSpans(t *testing.T) {
	got := windowSpans(3, 9, 3)
	if len(got) != 3 || got[0].Seq != 0 || got[0].From != 3 || got[0].To != 5 || got[0].Last != 5 || got[0].Lines != nil || got[2].From != 9 || got[2].To != 9 || got[2].Seq != 2 {
		t.Fatalf("%+v", got)
	}
	if windowSpans(6, 5, 3) != nil {
		t.Fatal("empty range")
	}
}
