//ff:func feature=port type=encoder control=sequence
//ff:what decimalRecord — a Wikidata decimal string ("+1.20", "-5", "1.5E+3") as an exact SIGNED-DECIMAL Value keeping its digits; false when it is not one
package port

import (
	"math/big"
	"regexp"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// wikidataDecimal is a quantity amount or bound: an optional sign, digits,
// an optional fraction and an optional exponent.
var wikidataDecimal = regexp.MustCompile(`^([+-]?)(\d+)(?:\.(\d+))?(?:[eE]([+-]?\d+))?$`)

// decimalRecord encodes s exactly as SIGNED-DECIMAL C×10^E: C is all the
// digits with the sign (so "+1.20" keeps its last zero: 120×10^-2), E minus
// the fraction length plus any exponent. It is false for anything else or an
// exponent outside int32.
func decimalRecord(s string) (core.Record, bool) {
	m := wikidataDecimal.FindStringSubmatch(s)
	if m == nil {
		return core.Record{}, false
	}
	c, _ := new(big.Int).SetString(m[1]+m[2]+m[3], 10)
	exp := int64(0)
	if m[4] != "" {
		var err error
		if exp, err = strconv.ParseInt(m[4], 10, 32); err != nil {
			return core.Record{}, false
		}
	}
	e := exp - int64(len(m[3]))
	if e < -1<<31 || e > 1<<31-1 {
		return core.Record{}, false
	}
	return core.EncodeSignedDecimalBig(int32(e), c), true
}
