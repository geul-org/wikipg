//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what TestPermOf — reordered properties, values swapped within a property, equal snaks in stored order, and no permutation when a snak has no row
package repair

import (
	"slices"
	"testing"
)

func TestPermOf(t *testing.T) {
	s := func(ord, vo int, p, v string) Snak {
		return Snak{Ordinal: ord, ValueOrdinal: vo, Property: p, SnakType: "value", Value: nullText(v)}
	}
	have := []Snak{s(0, 0, "P248", "Q1"), s(1, 0, "P813", "t"), s(2, 0, "P123", "Q7809"), s(3, 1, "P123", "Q2")}
	cases := []struct {
		want, have []Snak
		perm       []int
	}{
		{[]Snak{s(0, 0, "P123", "Q7809"), s(1, 1, "P123", "Q2"), s(2, 0, "P813", "t"), s(3, 0, "P248", "Q1")}, have, []int{2, 3, 1, 0}},
		{[]Snak{s(0, 0, "P248", "Q1"), s(1, 0, "P813", "t"), s(2, 0, "P123", "Q2"), s(3, 1, "P123", "Q7809")}, have, []int{0, 1, 3, 2}},
		{[]Snak{s(0, 0, "P1", "a"), s(1, 1, "P1", "a"), s(2, 0, "P2", "b")}, []Snak{s(0, 0, "P2", "b"), s(1, 0, "P1", "a"), s(2, 1, "P1", "a")}, []int{1, 2, 0}},
		{[]Snak{s(0, 0, "P248", "Q1"), s(1, 0, "P9", "x")}, have, nil},
	}
	for i, k := range cases {
		if got := permOf(k.want, k.have); !slices.Equal(got, k.perm) || (got == nil) != (k.perm == nil) {
			t.Fatalf("case %d: %v want %v", i, got, k.perm)
		}
	}
}
