//ff:func feature=port type=encoder control=selection
//ff:what quantityRecord — a Wikidata quantity as a QUANTITY Value: exact amount, Wikidata item unit or none, bounds or none; otherwise the reason
package port

import (
	"encoding/json"
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// quantityRecord converts a quantity datavalue (value/numeric/QUANTITY): the
// amount and bounds exactly (decimalRecord), the unit "1" as left out and an
// item unit as its WIKIDATA-QID, a missing bound left out (the Absent
// record). A value that is not JSON, has no or a malformed amount or bound,
// or a unit that is not a Wikidata item gives its reason (claimObject keeps
// it as original text).
func quantityRecord(value string) (core.Record, string) {
	var v struct {
		Amount string  `json:"amount"`
		Unit   string  `json:"unit"`
		Lower  *string `json:"lowerBound"`
		Upper  *string `json:"upperBound"`
	}
	absent := core.Record{Kind: core.KindReference}
	if err := json.Unmarshal([]byte(value), &v); err != nil {
		return core.Record{}, "quantity: JSON 아님"
	}
	q := core.Quantity{Unit: absent, Lower: absent, Upper: absent}
	amount, okAmount := decimalRecord(v.Amount)
	unit, okUnit := entityRef(strings.TrimPrefix(v.Unit, "http://www.wikidata.org/entity/"))
	bound := func(s *string, dst *core.Record) bool {
		if s == nil {
			return true
		}
		r, ok := decimalRecord(*s)
		*dst = r
		return ok
	}
	switch {
	case !okAmount:
		return core.Record{}, "quantity: 수"
	case v.Unit != "1" && (!okUnit || unit.Codec != core.RefWikidataQID<<3):
		return core.Record{}, "quantity: 단위"
	case !bound(v.Lower, &q.Lower) || !bound(v.Upper, &q.Upper):
		return core.Record{}, "quantity: 경계"
	}
	q.Amount = amount
	if v.Unit != "1" {
		q.Unit = unit
	}
	r, _ := core.EncodeQuantity(q) // the four records come from the core encoders, so encoding cannot fail
	return r, ""
}
