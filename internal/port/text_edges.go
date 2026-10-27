//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what textEdges — one LABEL, DESCRIPTION or ALIAS Edge per text, Edge Ids from next; texts with a malformed language counted as skipped
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikiport/internal/source"
)

// textEdges builds a LABEL, DESCRIPTION or ALIAS Edge (codec) on subject for each text,
// numbering Edge Ids from *next. A text whose language code is not a BCP 47
// tag in form is counted in rep.Skipped and gets no Edge.
func textEdges(codec uint8, a IDs, subject core.Record, texts []source.Text, rep *Report) []core.Record {
	var out []core.Record
	for _, t := range texts {
		r, ok := textEdge(codec, a, subject, t)
		if !ok {
			rep.Skipped["언어 코드: "+t.Lang]++
			continue
		}
		out = append(out, r)
	}
	return out
}
