//ff:type feature=source type=model
//ff:what Qualifier — one Wikidata qualifier snak as loaded into triple_qualifiers: claim, property, datatype, value
package source

// Qualifier is one row of geuldev.public.triple_qualifiers: the claim
// (triple ID) it qualifies, its property, the loader's datatype (time,
// wikibase-entityid, string, quantity, …) and the value as the loader stored
// it. Unlike a claim's entity value, an entity qualifier value is the
// datavalue JSON object ({"entity-type":…,"id":"Q5"}).
type Qualifier struct {
	TripleID int64
	Property string
	Type     string
	Value    string
}
