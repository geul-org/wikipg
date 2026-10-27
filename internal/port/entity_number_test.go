//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestEntityNumber — Q42 and P31 split; short, zero, signed, leading-zero and non-numeric IDs fail
package port

import "testing"

func TestEntityNumber(t *testing.T) {
	if p, n, ok := entityNumber("Q42"); !ok || p != 'Q' || n != 42 {
		t.Fatalf("Q42: %c %d %v", p, n, ok)
	}
	if p, n, ok := entityNumber("P31"); !ok || p != 'P' || n != 31 {
		t.Fatalf("P31: %c %d %v", p, n, ok)
	}
	for _, id := range []string{"", "Q", "Q0", "Q+5", "Q05", "Qx"} {
		if _, _, ok := entityNumber(id); ok {
			t.Fatalf("%q: want false", id)
		}
	}
}
