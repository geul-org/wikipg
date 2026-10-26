//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestDecimalRecord — signs, fractions with trailing zeros, exponents and 31 digits encode exactly; malformed strings, huge exponents and an exponent that overflows with the fraction fail
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecimalRecord(t *testing.T) {
	for s, want := range map[string]struct {
		e int32
		c string
	}{
		"+1.20": {-2, "120"}, "-5": {0, "-5"}, "+157.249": {-3, "157249"}, "1.5E+3": {2, "15"}, "+0.05": {-2, "5"},
		"+1234567890123456789012345678901": {0, "1234567890123456789012345678901"},
	} {
		r, ok := decimalRecord(s)
		e, c, err := core.DecodeSignedDecimalBig(r)
		if !ok || err != nil || e != want.e || c.String() != want.c {
			t.Fatalf("%s: %v %d %v %v", s, ok, e, c, err)
		}
	}
	for _, s := range []string{"", "+", "1.", "abc", "1e99999999999", "1.5e-2147483648"} {
		if _, ok := decimalRecord(s); ok {
			t.Fatalf("%q: want false", s)
		}
	}
}
