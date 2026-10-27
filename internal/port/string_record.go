//ff:func feature=port type=encoder control=selection
//ff:what stringRecord — a string datavalue by its datatype: url and Commons pages as URI, everything else as UTF8-TEXT
package port

import (
	"unicode/utf8"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/wikiport/internal/source"
)

// stringRecord converts a string datavalue by the property's datatype
// (files/geul/wikidata-datatype-mapping.md): url to a URI (urlRecord),
// commonsMedia to its Commons file page and geo-shape or tabular-data to its
// Commons Data page (commonsURI), and string, external-id, math,
// musical-notation and any other to UTF8-TEXT. A value that is not UTF-8 is
// skipped.
func stringRecord(c source.Claim) (core.Record, string) {
	switch {
	case !utf8.ValidString(c.Value):
		return core.Record{}, "문자열: UTF-8 아님"
	case c.Datatype == "url":
		return urlRecord(c.Value)
	case c.Datatype == "commonsMedia":
		return commonsURI("File:", c.Value), ""
	case c.Datatype == "geo-shape" || c.Datatype == "tabular-data":
		return commonsURI("", c.Value), ""
	default:
		return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(c.Value)), ""
	}
}
