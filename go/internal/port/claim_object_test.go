//ff:func feature=port type=encoder control=sequence
//ff:what TestClaimObject — converted values pass through, an unconvertible date becomes its original text with the reason, invalid UTF-8 and other skips stay skipped
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestClaimObject(t *testing.T) {
	if r, skip, raw := claimObject(source.Claim{Type: "wikibase-item", Value: "Q5"}); skip != "" || raw != "" || r.Kind != core.KindReference {
		t.Fatalf("item: %+v %q %q", r, skip, raw)
	}
	v := `{"time":"+1300-00-00T00:00:00Z","precision":7,"calendarmodel":"` + julian + `"}`
	r, skip, raw := claimObject(source.Claim{Type: "time", Value: v})
	b, err := core.DecodeBytesValue(r, core.CodecUTF8TextEven)
	if skip != "" || raw != "time: GEDCOM 대응 없는 정밀도 세기" || err != nil || string(b) != v {
		t.Fatalf("raw date: %q %q %v", skip, raw, err)
	}
	if _, skip, raw := claimObject(source.Claim{Type: "time", Value: "\xff"}); skip != "time: UTF-8 아님" || raw != "" {
		t.Fatalf("UTF-8: %q %q", skip, raw)
	}
	if _, skip, raw := claimObject(source.Claim{Type: "quantity", Value: "{}"}); skip != "형식 미지원: quantity" || raw != "" {
		t.Fatalf("quantity: %q %q", skip, raw)
	}
}
