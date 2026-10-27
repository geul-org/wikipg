//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestValidJulianDay — Julian leap years every fourth year (1100 yes, 1 BCE and 5 BCE yes, 2 BCE no); year 0 and bad months or days fail
package port

import "testing"

func TestValidJulianDay(t *testing.T) {
	cases := []struct {
		y    int64
		m, d int
		want bool
	}{
		{1100, 2, 29, true}, {1101, 2, 29, false}, {-1, 2, 29, true}, {-5, 2, 29, true}, {-2, 2, 29, false},
		{0, 1, 1, false}, {1066, 13, 1, false}, {1066, 4, 31, false}, {1066, 1, 0, false},
	}
	for _, c := range cases {
		if validJulianDay(c.y, c.m, c.d) != c.want {
			t.Fatalf("%+v", c)
		}
	}
}
