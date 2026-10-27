//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestDaysFromCivil — epoch, days around it, leap days and dates before year 0 match known day counts
package port

import "testing"

func TestDaysFromCivil(t *testing.T) {
	for _, c := range []struct {
		y    int64
		m, d int
		want int64
	}{
		{1970, 1, 1, 0},
		{1970, 1, 2, 1},
		{1969, 12, 31, -1},
		{2000, 2, 29, 11016},
		{1952, 3, 11, -6505},
		{1, 1, 1, -719162},
		{0, 3, 1, -719468},
		{-400, 3, 1, -719468 - 146097},
		{-1, 3, 1, -719468 - 366},
	} {
		if got := daysFromCivil(c.y, c.m, c.d); got != c.want {
			t.Fatalf("%d-%d-%d: %d want %d", c.y, c.m, c.d, got, c.want)
		}
	}
}
