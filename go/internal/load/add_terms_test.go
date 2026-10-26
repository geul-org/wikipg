//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddTerms — only kept languages are written; aliases carry their order
package load

import "testing"

func TestRowsAddTerms(t *testing.T) {
	r := NewRows()
	r.addTerms(Config{DumpID: 1, Langs: map[string]bool{"en": true}}, 3, sampleEntity())
	if r.Count("label") != 1 || r.Count("description") != 1 || r.Count("alias") != 2 {
		t.Fatalf("got %+v", r.byTable)
	}
	if a := r.byTable["alias"][1]; a[3] != 1 || a[4] != "Douglas Noël Adams" {
		t.Fatalf("alias %v", a)
	}
	ko := NewRows()
	ko.addTerms(Config{DumpID: 1, Langs: map[string]bool{"ko": true}}, 3, sampleEntity())
	if ko.Count("alias") != 0 || ko.Count("label") != 1 {
		t.Fatalf("ko %+v", ko.byTable)
	}
}
