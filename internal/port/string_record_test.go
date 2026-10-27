//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestStringRecord — url and Commons datatypes become URI Values, external-id and string stay UTF8-TEXT, bad UTF-8 and a schemeless url give reasons
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikiport/internal/source"
)

func TestStringRecord(t *testing.T) {
	for dt, codec := range map[string]uint8{"url": codecURIEven, "commonsMedia": codecURIEven, "geo-shape": codecURIEven, "tabular-data": codecURIEven, "external-id": core.CodecUTF8TextEven, "string": core.CodecUTF8TextEven, "math": core.CodecUTF8TextEven} {
		r, reason := stringRecord(source.Claim{Type: "string", Datatype: dt, Value: "http://x.org/a"})
		if reason != "" || r.Codec&^1 != codec {
			t.Fatalf("%s: codec 0x%02X %q", dt, r.Codec, reason)
		}
	}
	if _, reason := stringRecord(source.Claim{Datatype: "string", Value: "\xff"}); reason != "문자열: UTF-8 아님" {
		t.Fatalf("utf-8: %q", reason)
	}
	if _, reason := stringRecord(source.Claim{Datatype: "url", Value: "no-scheme"}); reason != "url: 스킴 없음" {
		t.Fatalf("scheme: %q", reason)
	}
}
