//ff:func feature=port type=helper control=sequence
//ff:what TestPropertiesOf — the predicate, plus the value when it is a property entity, not a P-like string
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestPropertiesOf(t *testing.T) {
	if got := propertiesOf(source.Claim{Property: "P31", Value: "Q5"}); !slices.Equal(got, []string{"P31"}) {
		t.Fatalf("item: %v", got)
	}
	if got := propertiesOf(source.Claim{Property: "P1659", Type: "wikibase-entityid", Value: "P279"}); !slices.Equal(got, []string{"P1659", "P279"}) {
		t.Fatalf("property: %v", got)
	}
	if got := propertiesOf(source.Claim{Property: "P214", Type: "external-id", Value: "P12"}); !slices.Equal(got, []string{"P214"}) {
		t.Fatalf("external ID: %v", got)
	}
}
