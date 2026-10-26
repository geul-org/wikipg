//ff:func feature=port type=encoder control=sequence
//ff:what claimObject — objectRecord, but a date, coordinate, monolingual text, url or quantity it cannot convert keeps its original datavalue as UTF8-TEXT and names why
package port

import (
	"strings"
	"unicode/utf8"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/wikiport/go/internal/source"
)

// claimObject returns the object record of c (objectRecord) and the reason it
// is skipped. Every date, coordinate, monolingual text, url and quantity is kept (wiki/README
// 값 변환): a time, globecoordinate, monolingualtext, url or quantity
// objectRecord cannot convert becomes its original
// datavalue JSON as UTF8-TEXT, and raw names the conversion reason so the
// report counts it.
func claimObject(c source.Claim) (obj core.Record, skip, raw string) {
	obj, skip = objectRecord(c)
	kept := strings.HasPrefix(skip, "time: ") || strings.HasPrefix(skip, "globecoordinate: ") || strings.HasPrefix(skip, "monolingualtext: ") || strings.HasPrefix(skip, "url: ") || strings.HasPrefix(skip, "quantity: ")
	if !kept {
		return obj, skip, ""
	}
	if !utf8.ValidString(c.Value) {
		return core.Record{}, strings.SplitN(skip, ":", 2)[0] + ": UTF-8 아님", ""
	}
	return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(c.Value)), "", skip
}
