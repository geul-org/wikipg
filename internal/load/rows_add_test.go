//ff:func feature=load type=model control=sequence
//ff:what TestRowsAdd — rows append to their table in order
package load

import "testing"

func TestRowsAdd(t *testing.T) {
	r := NewRows()
	r.add("label", "Q1", 1)
	r.add("label", "Q2", 1)
	if r.Count("label") != 2 || r.byTable["label"][1][0] != "Q2" {
		t.Fatalf("got %+v", r.byTable)
	}
}
