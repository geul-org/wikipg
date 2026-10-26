//ff:func feature=port type=encoder control=sequence
//ff:what textEdge — a LABEL, DESCRIPTION or ALIAS Edge: Edge Id, subject, the text as UTF8-LANGUAGE-TEXT; false for a malformed language code
package port

import (
	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/wikiport/go/internal/source"
)

// textEdge builds a LABEL (edge/LABEL), DESCRIPTION (edge/DESCRIPTION) or
// ALIAS (edge/ALIAS) Edge numbered id
// for t on subject: the Object is t as UTF8-LANGUAGE-TEXT (languageText). It
// is false when t's language code is not a BCP 47 tag in form.
func textEdge(codec uint8, a IDs, subject core.Record, t source.Text) (core.Record, bool) {
	text, ok := languageText(t.Lang, t.Value)
	if !ok {
		return core.Record{}, false
	}
	r, _ := edgeWith(a, codec, []core.Record{subject, text})
	return r, true
}
