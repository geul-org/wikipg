//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what TestTallyAdd — each verdict counts in its own field
package repair

import "testing"

func TestTallyAdd(t *testing.T) {
	var tl Tally
	for _, v := range []Verdict{Same, Order, Order, Content, Missing, Missing, Missing} {
		tl.add(v)
	}
	if tl != (Tally{Same: 1, Order: 2, Content: 1, Missing: 3}) {
		t.Fatalf("%+v", tl)
	}
}
