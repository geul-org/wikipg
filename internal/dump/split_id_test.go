//ff:func feature=dump type=decoder control=iteration dimension=1
//ff:what TestSplitID — Q, P and L IDs split; short, zero-led, signed and non-numeric IDs are false
package dump

import "testing"

func TestSplitID(t *testing.T) {
	if k, n, ok := SplitID("Q42"); !ok || k != 'Q' || n != 42 {
		t.Fatalf("Q42: %c %d %v", k, n, ok)
	}
	if k, n, ok := SplitID("P31"); !ok || k != 'P' || n != 31 {
		t.Fatalf("P31: %c %d %v", k, n, ok)
	}
	for _, id := range []string{"Q", "Q0", "Q01", "Q+1", "Qx", "Q99999999999999999999"} {
		if _, _, ok := SplitID(id); ok {
			t.Fatalf("%q: want false", id)
		}
	}
}
