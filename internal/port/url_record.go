//ff:func feature=port type=encoder control=sequence
//ff:what urlRecord — a Wikidata url as a URI Value, non-ASCII and disallowed bytes percent-encoded; a reason when there is no scheme
package port

import (
	"regexp"

	"github.com/geul-org/gogeul/pkg/core"
)

// uriScheme is a URI scheme followed by its colon (RFC 3986 3.1).
var uriScheme = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9+.-]*:`)

// urlRecord converts a url value (an IRI as Wikidata stores it) to a URI
// Value (0x44/0x45): every byte a URI cannot hold, including each UTF-8 byte
// of a non-ASCII character, is percent-encoded (RFC 3987 3.1), which a
// reader can undo. A value with no scheme is not a URI and gives its reason
// (claimObject keeps it as original text).
func urlRecord(s string) (core.Record, string) {
	u := percentEncode(s, isURIByte)
	if !uriScheme.MatchString(u) {
		return core.Record{}, "url: 스킴 없음"
	}
	return core.EncodeBytesValue(codecURIEven, []byte(u)), ""
}
