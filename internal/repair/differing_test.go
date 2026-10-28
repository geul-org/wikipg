//ff:func feature=repair type=helper control=sequence
//ff:what TestDiffering — every verdict but same is to be replaced
package repair

import "testing"

func TestDiffering(t *testing.T) {
	got := differing([]HashCheck{{Hash: "s"}, {Hash: "o", Verdict: Order}, {Hash: "c", Verdict: Content}, {Hash: "m", Verdict: Missing}})
	if len(got) != 3 || got[0].Hash != "o" || got[2].Hash != "m" {
		t.Fatalf("%v", got)
	}
}
