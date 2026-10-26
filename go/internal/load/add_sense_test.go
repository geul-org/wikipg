//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddSense — entity row of type sense, sense row, glosses in kept languages only, statements; key range ends with false
package load

import "testing"

func TestRowsAddSense(t *testing.T) {
	r := NewRows()
	s := sampleLexeme().Senses[0]
	next, ok := r.addSense(Config{DumpID: 1, Langs: map[string]bool{"fr": true}}, 3, "L4", 4, 0, s, 2, NewRefSet(64))
	if !ok || next != 3 || r.Count("entity") != 1 || r.Count("sense") != 1 || r.Count("sense_gloss") != 1 || r.byTable["sense_gloss"][0][2] != "fr" || r.Count("statement") != 1 {
		t.Fatalf("got %d %v %+v", next, ok, r.byTable)
	}
	if _, ok := NewRows().addSense(Config{DumpID: 0}, 3, "L4", 4, 0, s, 0, NewRefSet(64)); ok {
		t.Fatal("bad dump: want false")
	}
}
