//ff:func feature=port type=encoder control=sequence
//ff:what tripleEdge — a TRIPLE Edge: Edge Id, Predicate, Subject, Object, rank word
package port

import "github.com/geul-org/geul/go/pkg/core"

// tripleEdge builds a TRIPLE Edge (edge/TRIPLE): the LOCAL Edge Id, the
// predicate (WIKIDATA-PID), the subject, the object record and the rank as
// one raw word (rankWord).
func tripleEdge(id uint64, predicate, subject, object core.Record, rank core.Word) core.Record {
	r := compose(core.KindEdge, edgeTriple, core.EncodeLocalRef(id), predicate, subject, object)
	r.Payload = append(r.Payload, rank)
	return r
}
