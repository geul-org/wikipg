//ff:func feature=port type=encoder control=sequence
//ff:what tripleEdge — a TRIPLE Edge: Edge Id, Predicate, Subject, Object, rank word
package port

import "github.com/geul-org/gogeul/pkg/core"

// tripleEdge builds a TRIPLE Edge (edge/TRIPLE): the LOCAL Edge Id, the
// predicate (WIKIDATA-PID), the subject, the object record and the rank as
// one raw word (rankWord).
func tripleEdge(a IDs, predicate, subject, object core.Record, rank core.Word) (core.Record, uint64) {
	return edgeWith(a, edgeTriple, []core.Record{predicate, subject, object}, rank)
}
