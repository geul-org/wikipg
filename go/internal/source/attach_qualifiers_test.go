//ff:func feature=source type=loader control=sequence
//ff:what TestAttachQualifiers — qualifiers go to their claim in order; one of an unlisted claim is dropped
package source

import "testing"

func TestAttachQualifiers(t *testing.T) {
	cs := []Claim{{ID: 1}, {ID: 2}}
	attachQualifiers(cs, []Qualifier{{TripleID: 2, Property: "P580"}, {TripleID: 9, Property: "P1"}, {TripleID: 2, Property: "P582"}})
	if len(cs[0].Qualifiers) != 0 || len(cs[1].Qualifiers) != 2 || cs[1].Qualifiers[0].Property != "P580" || cs[1].Qualifiers[1].Property != "P582" {
		t.Fatalf("got %+v", cs)
	}
}
