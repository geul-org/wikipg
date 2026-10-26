//ff:func feature=port type=encoder control=sequence
//ff:what qualifierEdge — a QUALIFIER Edge: Edge Id, Property, Target Edge Id, Value
package port

import "github.com/geul-org/geul/go/pkg/core"

// qualifierEdge builds a QUALIFIER Edge (edge/QUALIFIER) numbered id that
// qualifies the Edge target: the LOCAL Edge Id, the property (WIKIDATA-PID),
// the target as a LOCAL Reference and the value record.
func qualifierEdge(id uint64, property core.Record, target uint64, value core.Record) core.Record {
	return compose(core.KindEdge, edgeQualifier, core.EncodeLocalRef(id), property, core.EncodeLocalRef(target), value)
}
