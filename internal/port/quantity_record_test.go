//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestQuantityRecord — 652 km ±1 and a unitless 1.20 become QUANTITY with their parts; bad JSON, amount, unit and bound give reasons
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQuantityRecord(t *testing.T) {
	r, reason := quantityRecord(`{"amount":"+652","unit":"http://www.wikidata.org/entity/Q828224","lowerBound":"+651","upperBound":"+653"}`)
	q, err := core.DecodeQuantity(r)
	if reason != "" || err != nil || core.IsAbsent(q.Unit) || core.IsAbsent(q.Lower) || core.IsAbsent(q.Upper) {
		t.Fatalf("km: %q %v %+v", reason, err, q)
	}
	r, reason = quantityRecord(`{"amount":"+1.20","unit":"1"}`)
	q, err = core.DecodeQuantity(r)
	if reason != "" || err != nil || !core.IsAbsent(q.Unit) || !core.IsAbsent(q.Lower) || !core.IsAbsent(q.Upper) {
		t.Fatalf("unitless: %q %v", reason, err)
	}
	if e, c, _ := core.DecodeSignedDecimalBig(q.Amount); e != -2 || c.Int64() != 120 {
		t.Fatalf("quantum: %d %v", e, c)
	}
	for want, v := range map[string]string{
		"quantity: JSON 아님": "{",
		"quantity: 수":       `{"amount":"x","unit":"1"}`,
		"quantity: 단위":      `{"amount":"+1","unit":"http://example.org/m"}`,
		"quantity: 경계":      `{"amount":"+1","unit":"1","lowerBound":"?"}`,
	} {
		if _, got := quantityRecord(v); got != want {
			t.Fatalf("%s: %q want %q", v, got, want)
		}
	}
}
