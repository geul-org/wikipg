//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what TestCompareSnaks — same rows, nothing stored, reordered, a doubled snak for a missing one, and a shorter list
package repair

import "testing"

func TestCompareSnaks(t *testing.T) {
	a, b, c := Snak{Ordinal: 0, Property: "P248"}, Snak{Ordinal: 1, Property: "P813"}, Snak{Ordinal: 2, Property: "P123"}
	ba, ab := Snak{Ordinal: 0, Property: "P813"}, Snak{Ordinal: 1, Property: "P248"}
	cases := []struct {
		want, have []Snak
		v          Verdict
	}{
		{[]Snak{a, b}, []Snak{a, b}, Same},
		{nil, nil, Same},
		{[]Snak{a}, nil, Missing},
		{[]Snak{a, b}, []Snak{ba, ab}, Order},
		{[]Snak{a, b, c}, []Snak{a, b, {Ordinal: 2, Property: "P813"}}, Content},
		{[]Snak{a, b}, []Snak{a}, Content},
	}
	for i, k := range cases {
		if got := compareSnaks(k.want, k.have); got != k.v {
			t.Fatalf("case %d: %v want %v", i, got, k.v)
		}
	}
}
