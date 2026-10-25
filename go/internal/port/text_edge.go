//ff:func feature=port type=encoder control=sequence
//ff:what textEdge — a LABEL or ALIAS Edge: Edge Id, LANGUAGE-BCP47-CODE, Subject, UTF8-NFC-TEXT
package port

import (
	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// textEdge builds a LABEL (edge/LABEL) or ALIAS (proposal, same layout) Edge
// for t on subject: the language as LANGUAGE-BCP47-CODE (the GEUL language
// registry is not sealed, so no LANGUAGE-SHORT-CODE is assigned) and the text
// as UTF8-NFC-TEXT (Wikibase stores NFC).
func textEdge(codec uint8, id uint64, subject core.Record, t source.Text) core.Record {
	lang := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(t.Lang))
	text := core.EncodeBytesValue(codecUTF8NFCTextEven, []byte(t.Value))
	return compose(core.KindEdge, codec, core.EncodeLocalRef(id), lang, subject, text)
}
