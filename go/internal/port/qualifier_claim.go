//ff:func feature=port type=helper control=sequence
//ff:what qualifierClaim — a qualifier as a claim value claimObject converts
package port

import "github.com/geul-org/wikiport/go/internal/source"

// qualifierClaim returns q as a claim (property, type, datatype and value) so the value
// converts like a main value (claimObject). Both come in the same form from
// the loaded dump (source.queriesFor: an entity is a bare ID).
func qualifierClaim(q source.Qualifier) source.Claim {
	return source.Claim{Property: q.Property, Type: q.Type, Datatype: q.Datatype, Value: q.Value}
}
