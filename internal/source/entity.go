//ff:type feature=source type=model
//ff:what Entity — one Wikidata item or property with its labels, descriptions, aliases and claims in the chosen languages
package source

// Entity is one Wikidata item (Q) or property (P) of a loaded dump: a
// property's datatype (url, external-id …; empty for an item),
// labels, descriptions and aliases in the chosen languages (aliases in alias_order) and
// every claim in statement ID order.
type Entity struct {
	ID           string
	Datatype     string
	Labels       []Text
	Descriptions []Text
	Aliases      []Text
	Claims       []Claim
}
