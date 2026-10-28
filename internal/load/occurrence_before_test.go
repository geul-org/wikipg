//ff:func feature=load type=helper control=sequence
//ff:what TestOccurrenceBefore — statement key decides first, then the reference position; equal is not before
package load

import "testing"

func TestOccurrenceBefore(t *testing.T) {
	a, b, c := Occurrence{Statement: 1, Ordinal: 5}, Occurrence{Statement: 2, Ordinal: 0}, Occurrence{Statement: 2, Ordinal: 1}
	if !a.Before(b) || b.Before(a) || !b.Before(c) || c.Before(b) || b.Before(b) {
		t.Fatal("order")
	}
}
