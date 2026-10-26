//ff:func feature=port type=encoder control=sequence
//ff:what commonsURI — a Commons page name as a URI Value: https://commons.wikimedia.org/wiki/ + prefix + the name with _ for spaces, one encoded segment
package port

import (
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// commonsURI points at the Wikimedia Commons page prefix+name as a URI Value
// (wiki/README 값 변환): a commonsMedia file name with prefix "File:", a
// geo-shape or tabular-data page name ("Data:…") with none. Spaces become _
// as in MediaWiki titles and the title is percent-encoded as one path
// segment (isSegmentByte), so ?, # and / in a title stay part of it.
func commonsURI(prefix, name string) core.Record {
	title := percentEncode(prefix+strings.ReplaceAll(name, " ", "_"), isSegmentByte)
	return core.EncodeBytesValue(codecURIEven, []byte("https://commons.wikimedia.org/wiki/"+title))
}
