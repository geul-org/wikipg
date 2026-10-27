//ff:func feature=load type=model control=iteration dimension=1
//ff:what TestRowsAddLexeme — lexeme, lemmas (unfiltered), form with features and representation, sense with kept glosses, entity rows for form and sense, statement keys after the lexeme's; an item adds nothing
package load

import "testing"

func TestRowsAddLexeme(t *testing.T) {
	r := NewRows()
	c := Config{DumpID: 1, Langs: map[string]bool{"en": true}}
	r.addLexeme(c, 3, sampleLexeme(), 4, 1, NewRefSet(64))
	for table, want := range map[string]int{"lexeme": 1, "lemma": 2, "form": 1, "form_representation": 1, "sense": 1, "sense_gloss": 1, "entity": 2, "statement": 2} {
		if got := r.Count(table); got != want {
			t.Fatalf("%s: %d want %d", table, got, want)
		}
	}
	st := r.byTable["statement"]
	if st[0][2] != "L4-F1" || st[0][0] != int64(1<<55|3<<20|1) || st[1][2] != "L4-S1" || st[1][0] != int64(1<<55|3<<20|2) {
		t.Fatalf("statements %v", st)
	}
	if e := r.byTable["entity"]; e[0][0] != "L4-F1" || e[0][4] != "form" || e[1][4] != "sense" || e[0][3] != int64(4) {
		t.Fatalf("entity rows %v", e)
	}
	item := NewRows()
	item.addLexeme(c, 3, sampleEntity(), 42, 0, NewRefSet(64))
	if item.Count("lexeme") != 0 {
		t.Fatal("item: want nothing")
	}
}
