//ff:func feature=port type=encoder control=sequence
//ff:what statementEdge — a STATEMENT Edge: Edge Id, Predicate, Subject, Object and the rank word
package port

import "github.com/geul-org/geul/go/pkg/core"

// statementEdge builds a STATEMENT Edge (proposal layout): the LOCAL Edge Id,
// the predicate (WIKIDATA-PID), the subject, the object record and the rank
// as one raw word (rankWord).
func statementEdge(id uint64, predicate, subject, object core.Record, rank core.Word) core.Record {
	r := compose(core.KindEdge, edgeStatement, core.EncodeLocalRef(id), predicate, subject, object)
	r.Payload = append(r.Payload, rank)
	return r
}
