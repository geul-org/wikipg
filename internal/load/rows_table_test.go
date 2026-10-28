//ff:func feature=load type=model control=sequence
//ff:what TestRowsTable — a table's rows in the order added; none for a table without rows
package load

import "testing"

func TestRowsTable(t *testing.T) {
	r := NewRows()
	r.add("reference", "a", 1)
	r.add("reference", "b", 1)
	if got := r.Table("reference"); len(got) != 2 || got[1][0] != "b" {
		t.Fatalf("got %v", got)
	}
	if r.Table("statement") != nil {
		t.Fatal("empty table")
	}
}
