//ff:func feature=port type=encoder control=selection
//ff:what objectRecord — the GEUL object record of a claim value (entity, string, time, monolingual text, coordinate); otherwise the reason the claim is skipped
package port

import (
	"unicode/utf8"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// objectRecord converts the value of c: an entity ID to its Reference
// (entityRef), a string or external ID to UTF8-TEXT, a time to GREGORIAN-DATE,
// EDTF or GEDCOM-DATE (timeRecord), a monolingualtext to UTF8-LANGUAGE-TEXT
// (monolingualRecord), a globecoordinate to COORDINATE-EARTH or -GEO
// (coordinateRecord). A time or coordinate that cannot be converted has a
// "time: …" or "globecoordinate: …" reason; claimObject keeps it as original
// text. An empty value is a novalue
// or somevalue snak. Other value
// types are not converted in this pass (wiki/README 작업 범위); the reason
// names what was skipped.
func objectRecord(c source.Claim) (core.Record, string) {
	if c.Value == "" {
		return core.Record{}, "값 없음(novalue·somevalue)"
	}
	switch c.Type {
	case "wikibase-entityid", "wikibase-item", "wikibase-property":
		if r, ok := entityRef(c.Value); ok {
			return r, ""
		}
		return core.Record{}, "개체: " + c.Value[:1]
	case "string", "external-id":
		if !utf8.ValidString(c.Value) {
			return core.Record{}, "문자열: UTF-8 아님"
		}
		return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(c.Value)), ""
	case "time":
		return timeRecord(c.Value)
	case "monolingualtext":
		return monolingualRecord(c.Value)
	case "globecoordinate", "globe-coordinate":
		return coordinateRecord(c.Value)
	default:
		return core.Record{}, "형식 미지원: " + c.Type
	}
}
