//ff:type feature=source type=model
//ff:what Claim — one Wikidata main snak as loaded into triples: property, value type, value, rank, and its qualifiers
package source

// Claim is one row of geuldev.public.triples: the property of the main snak,
// the loader's value type (a datavalue type such as wikibase-entityid, string,
// time, or a datatype such as wikibase-item), the value as the loader stored
// it (an entity ID, a string or a datavalue JSON object; empty for novalue and
// somevalue), the rank (preferred, normal, deprecated) and its qualifiers in
// load order.
type Claim struct {
	ID         int64
	Property   string
	Type       string
	Value      string
	Rank       string
	Qualifiers []Qualifier
}
