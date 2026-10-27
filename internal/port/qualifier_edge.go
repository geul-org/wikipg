//ff:func feature=port type=encoder control=sequence
//ff:what qualifierEdge — a QUALIFIER Edge: Edge Id, Property, Target Edge Id, Value
package port

import "github.com/geul-org/gogeul/pkg/core"

// qualifierEdge builds a QUALIFIER Edge (edge/QUALIFIER) numbered id that
// qualifies the Edge target: the LOCAL Edge Id, the property (WIKIDATA-PID),
// the target as a LOCAL Reference and the value record.
func qualifierEdge(a IDs, property core.Record, target uint64, value core.Record) core.Record {
	r, _ := edgeWith(a, edgeQualifier, []core.Record{property, core.EncodeLocalRef(target), value})
	return r
}
