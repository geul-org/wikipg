//ff:func feature=port type=helper control=sequence
//ff:what propertiesOf — the property IDs a converted claim uses: its predicate, and its value when that is a property
package port

import (
	"strings"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// propertiesOf returns the properties a converted claim uses: the predicate,
// and the value when it is a property ID (P…).
func propertiesOf(c source.Claim) []string {
	ps := []string{c.Property}
	if strings.HasPrefix(c.Value, "P") {
		ps = append(ps, c.Value)
	}
	return ps
}
