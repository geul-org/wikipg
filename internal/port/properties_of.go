//ff:func feature=port type=helper control=sequence
//ff:what propertiesOf — the property IDs a converted claim uses: its predicate, and its value when that is a property
package port

import (
	"strings"

	"github.com/geul-org/wikiport/internal/source"
)

// propertiesOf returns the properties a converted claim uses: the predicate,
// and the value when it is a property entity (P…). A string or external ID
// that happens to start with P is not a property.
func propertiesOf(c source.Claim) []string {
	ps := []string{c.Property}
	entity := c.Type == "wikibase-entityid" || c.Type == "wikibase-item" || c.Type == "wikibase-property"
	if entity && strings.HasPrefix(c.Value, "P") {
		ps = append(ps, c.Value)
	}
	return ps
}
