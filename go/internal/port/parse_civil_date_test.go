//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestParseCivilDate — CE, BCE and coarse-precision timestamps parse; malformed ones fail
package port

import "testing"

func TestParseCivilDate(t *testing.T) {
	if y, m, d, ok := parseCivilDate("+1952-03-11T00:00:00Z"); !ok || y != 1952 || m != 3 || d != 11 {
		t.Fatalf("CE: %d %d %d %v", y, m, d, ok)
	}
	if y, _, _, ok := parseCivilDate("-0044-03-15T00:00:00Z"); !ok || y != -44 {
		t.Fatalf("BCE: %d %v", y, ok)
	}
	if y, m, d, ok := parseCivilDate("+1974-00-00T00:00:00Z"); !ok || y != 1974 || m != 0 || d != 0 {
		t.Fatalf("year: %d %d %d %v", y, m, d, ok)
	}
	for _, s := range []string{"", "1952-03-11T00:00:00Z", "+1952-03-11T01:00:00Z", "+1952-03T00:00:00Z", "+1952-3-11T00:00:00Z", "+1952-03-1T00:00:00Z", "+x952-03-11T00:00:00Z", "+1952-0x-11T00:00:00Z", "+1952-03-1xT00:00:00Z"} {
		if _, _, _, ok := parseCivilDate(s); ok {
			t.Fatalf("%q: want false", s)
		}
	}
}
