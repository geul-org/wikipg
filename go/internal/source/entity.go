//ff:type feature=source type=model
//ff:what Entity — one Wikidata item or property with its labels, aliases and claims in the chosen languages
package source

// Entity is one Wikidata item (Q) or property (P) as loaded from geuldev:
// labels and aliases in the chosen languages (aliases in alias_order) and
// every claim in triple ID order.
type Entity struct {
	ID      string
	Labels  []Text
	Aliases []Text
	Claims  []Claim
}
