//ff:func feature=port type=helper control=sequence
//ff:what qualifierClaim — a qualifier as a claim value claimObject converts (an entity datavalue JSON becomes its ID)
package port

import (
	"encoding/json"
	"strings"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// qualifierClaim returns q as a claim (property, type and value) so the value
// converts like a main value (claimObject). An entity qualifier value is the
// datavalue JSON object, while a claim's entity value is the bare ID, so its
// "id" is taken; a value that does not parse is kept and fails as an entity.
func qualifierClaim(q source.Qualifier) source.Claim {
	c := source.Claim{Property: q.Property, Type: q.Type, Value: q.Value}
	var v struct {
		ID string `json:"id"`
	}
	if q.Type == "wikibase-entityid" && strings.HasPrefix(q.Value, "{") && json.Unmarshal([]byte(q.Value), &v) == nil && v.ID != "" {
		c.Value = v.ID
	}
	return c
}
