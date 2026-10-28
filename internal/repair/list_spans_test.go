//ff:func feature=repair type=helper control=sequence
//ff:what TestListSpans — lines cut into numbered batches, each checkpointed at its last line; none gives none
package repair

import "testing"

func TestListSpans(t *testing.T) {
	got := listSpans([]int64{2, 7, 9, 40, 41}, 2)
	if len(got) != 3 || got[0].Last != 7 || len(got[1].Lines) != 2 || got[2].Seq != 2 || got[2].Last != 41 || got[2].Lines[0] != 41 {
		t.Fatalf("%+v", got)
	}
	if listSpans(nil, 2) != nil {
		t.Fatal("none")
	}
}
