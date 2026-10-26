//ff:func feature=load type=model control=sequence
//ff:what TestRowsCount — counts per table; an unknown table has none
package load

import "testing"

func TestRowsCount(t *testing.T) {
	r := NewRows()
	r.add("alias", "Q1")
	if r.Count("alias") != 1 || r.Count("nope") != 0 {
		t.Fatal("wrong count")
	}
}
