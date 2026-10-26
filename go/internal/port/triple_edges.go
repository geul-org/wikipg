//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what tripleEdges — a TRIPLE and then its QUALIFIERs per convertible claim, Edge Ids from next; skipped claims counted by reason
package port

import (
	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// tripleEdges builds a TRIPLE on subject for each claim whose predicate,
// value and rank convert, each followed by its QUALIFIERs (qualifierEdges),
// numbering Edge Ids from *next. Every other claim is
// counted in rep.Skipped by reason, a date kept as original text in rep.Raw
// by its conversion reason (claimObject), and every property used as a predicate or
// a value is added to rep.Properties.
func tripleEdges(next *uint64, subject core.Record, claims []source.Claim, rep *Report) []core.Record {
	var out []core.Record
	for _, c := range claims {
		pred, okPred := entityRef(c.Property)
		obj, reason, raw := claimObject(c)
		rank, okRank := rankWord(c.Rank)
		switch {
		case !okPred:
			reason = "속성 ID: " + c.Property
		case reason == "" && !okRank:
			reason = "rank: " + c.Rank
		}
		if reason != "" {
			rep.Skipped[reason]++
			continue
		}
		if raw != "" {
			rep.Raw[raw]++
		}
		rep.Properties = append(rep.Properties, propertiesOf(c)...)
		id := *next
		out = append(out, tripleEdge(id, pred, subject, obj, rank))
		*next++
		rep.Triples++
		out = append(out, qualifierEdges(next, id, c.Qualifiers, rep)...)
	}
	return out
}
