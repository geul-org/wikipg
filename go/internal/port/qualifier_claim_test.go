//ff:func feature=port type=helper control=sequence
//ff:what TestQualifierClaim — an entity datavalue JSON becomes its ID; other types and unparsable entity values stay as they are
package port

import (
	"testing"

	"github.com/geul-org/wikiport/go/internal/source"
)

func TestQualifierClaim(t *testing.T) {
	if c := qualifierClaim(source.Qualifier{Property: "P512", Type: "wikibase-entityid", Value: `{"entity-type":"item","numeric-id":1765120,"id":"Q1765120"}`}); c.Value != "Q1765120" || c.Property != "P512" {
		t.Fatalf("entity: %+v", c)
	}
	if c := qualifierClaim(source.Qualifier{Type: "string", Value: "{x"}); c.Value != "{x" {
		t.Fatalf("string: %+v", c)
	}
	if c := qualifierClaim(source.Qualifier{Type: "wikibase-entityid", Value: "{bad"}); c.Value != "{bad" {
		t.Fatalf("bad JSON: %+v", c)
	}
}
