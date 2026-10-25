//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestGedcomValid — Julian dates pass; unparsable values, a dual year that is not the next year, non-Julian calendars fail
package port

import "testing"

func TestGedcomValid(t *testing.T) {
	for _, s := range []string{"JULIAN 14 OCT 1066", "JULIAN 1300", "JULIAN 15 MAR 44 BCE"} {
		if err := gedcomValid(s); err != nil {
			t.Fatalf("%s: %v", s, err)
		}
	}
	for _, s := range []string{"JULIAN 13 1300", "JULIAN OCT", "FOO 1300", "JULIAN 21 FEB 1750/53", "14 OCT 1066", ""} {
		if gedcomValid(s) == nil {
			t.Fatalf("%q: want error", s)
		}
	}
}
