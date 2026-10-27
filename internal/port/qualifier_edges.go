//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what qualifierEdges — a QUALIFIER on a TRIPLE for each convertible qualifier; others counted by reason
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikiport/internal/source"
)

// qualifierEdges builds a QUALIFIER on the Edge target for each qualifier in
// qs whose property and value convert (qualifierClaim, claimObject),
// numbering Edge Ids from *next and keeping qs order. Skipped qualifiers are
// counted in rep.Skipped and dates kept as original text in rep.Raw, both
// prefixed "qualifier "; the properties used are added to rep.Properties.
func qualifierEdges(a IDs, target uint64, qs []source.Qualifier, rep *Report) []core.Record {
	var out []core.Record
	for _, q := range qs {
		c := qualifierClaim(q)
		prop, okProp := entityRef(q.Property)
		value, reason, raw := claimObject(c)
		if !okProp {
			reason = "속성 ID: " + q.Property
		}
		if reason != "" {
			rep.Skipped["qualifier "+reason]++
			continue
		}
		if raw != "" {
			rep.Raw["qualifier "+raw]++
		}
		rep.Properties = append(rep.Properties, propertiesOf(c)...)
		out = append(out, qualifierEdge(a, prop, target, value))
		rep.Qualifiers++
	}
	return out
}
