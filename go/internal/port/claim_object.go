//ff:func feature=port type=encoder control=sequence
//ff:what claimObject — objectRecord, but a date it cannot convert keeps its original datavalue as UTF8-TEXT and names why
package port

import (
	"strings"
	"unicode/utf8"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// claimObject returns the object record of c (objectRecord) and the reason it
// is skipped. Every date is kept (wiki/README 값 변환): a time
// objectRecord cannot convert becomes its original datavalue JSON as
// UTF8-TEXT, and raw names the conversion reason so the report counts it.
func claimObject(c source.Claim) (obj core.Record, skip, raw string) {
	obj, skip = objectRecord(c)
	if c.Type != "time" || !strings.HasPrefix(skip, "time: ") {
		return obj, skip, ""
	}
	if !utf8.ValidString(c.Value) {
		return core.Record{}, "time: UTF-8 아님", ""
	}
	return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(c.Value)), "", skip
}
