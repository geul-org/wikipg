//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestValidDay — real days pass; month 0, 13, day 0, 31 April and 29 February outside leap years fail
package port

import "testing"

func TestValidDay(t *testing.T) {
	for _, c := range []struct {
		y    int64
		m, d int
		want bool
	}{
		{2024, 2, 29, true},
		{2000, 2, 29, true},
		{1900, 2, 29, false},
		{2023, 2, 29, false},
		{2023, 4, 30, true},
		{2023, 4, 31, false},
		{2023, 0, 1, false},
		{2023, 13, 1, false},
		{2023, 1, 0, false},
	} {
		if validDay(c.y, c.m, c.d) != c.want {
			t.Fatalf("%d-%d-%d: want %v", c.y, c.m, c.d, c.want)
		}
	}
}
