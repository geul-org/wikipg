//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what textEdges — one LABEL or ALIAS Edge per text, Edge Ids from next
package port

import (
	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// textEdges builds a LABEL or ALIAS Edge (codec) on subject for each text,
// numbering Edge Ids from *next.
func textEdges(codec uint8, next *uint64, subject core.Record, texts []source.Text) []core.Record {
	var out []core.Record
	for _, t := range texts {
		out = append(out, textEdge(codec, *next, subject, t))
		*next++
	}
	return out
}
