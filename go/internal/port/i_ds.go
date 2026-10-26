//ff:type feature=port type=model
//ff:what IDs — where a ported document gets its LOCAL IDs: the GEUL library's carried-over assignment (delta.LocalIDs)
package port

import "github.com/geul-org/geul/go/pkg/core"

// IDs hands out LOCAL IDs that carry over from the previous version of the
// document (OVERVIEW 개정: same statement same ID, new ones from the auto
// increment). *delta.LocalIDs of the GEUL library implements it.
type IDs interface {
	Edge(codec uint8, rest []core.Word) uint64
	Define(object core.Record) (edge, local uint64)
}
