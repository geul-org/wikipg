//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestEdtfValid — real dates and reduced precisions pass; impossible days, months, Y-prefixed years and junk fail
package port

import "testing"

func TestEdtfValid(t *testing.T) {
	for _, s := range []string{"1952-03-11", "1974", "197X", "-0043-03-15"} {
		if err := edtfValid(s); err != nil {
			t.Fatalf("%s: %v", s, err)
		}
	}
	for _, s := range []string{"2023-02-29", "1974-13", "Y-13799999999", "abc"} {
		if err := edtfValid(s); err == nil {
			t.Fatalf("%s: want error", s)
		}
	}
}
