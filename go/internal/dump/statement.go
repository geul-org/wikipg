//ff:type feature=dump type=model
//ff:what Statement — one Wikidata claim: GUID, rank, main snak, qualifiers with their order, references
package dump

// Statement is one claim: its GUID (Q42$…), rank, main snak, qualifiers by
// property with the property order (qualifiers-order) and references in order.
type Statement struct {
	ID              string            `json:"id"`
	Rank            string            `json:"rank"`
	Mainsnak        Snak              `json:"mainsnak"`
	Qualifiers      map[string][]Snak `json:"qualifiers"`
	QualifiersOrder []string          `json:"qualifiers-order"`
	References      []Reference       `json:"references"`
}
