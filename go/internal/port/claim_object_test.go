//ff:func feature=port type=encoder control=sequence
//ff:what TestClaimObject — converted values pass through, an unconvertible date, coordinate or monolingual text becomes its original text with the reason, invalid UTF-8 and other skips stay skipped; a schemeless url and a malformed quantity are kept as text
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/wikiport/go/internal/source"
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
	if _, skip, raw := claimObject(source.Claim{Type: "globecoordinate", Value: `{"latitude":1,"longitude":2,"globe":""}`}); skip != "" || raw != "globecoordinate: 천체 없음" {
		t.Fatalf("raw coordinate: %q %q", skip, raw)
	}
	if _, skip, raw := claimObject(source.Claim{Type: "monolingualtext", Value: `{"text":"x","language":""}`}); skip != "" || raw != "monolingualtext: 언어 코드 " {
		t.Fatalf("raw monolingual: %q %q", skip, raw)
	}
	if r, skip, raw := claimObject(source.Claim{Type: "string", Datatype: "url", Value: "example.org"}); skip != "" || raw != "url: 스킴 없음" || r.Codec&^1 != core.CodecUTF8TextEven {
		t.Fatalf("raw url: %q %q", skip, raw)
	}
	if r, skip, raw := claimObject(source.Claim{Type: "quantity", Value: `{"amount":"x","unit":"1"}`}); skip != "" || raw != "quantity: 수" || r.Codec&^1 != core.CodecUTF8TextEven {
		t.Fatalf("raw quantity: %q %q", skip, raw)
	}
	if _, skip, raw := claimObject(source.Claim{Type: "unknown", Value: "{}"}); skip != "형식 미지원: unknown" || raw != "" {
		t.Fatalf("unknown type: %q %q", skip, raw)
	}
}
