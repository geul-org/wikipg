//ff:type feature=dump type=model
//ff:what Reference — one source of a claim: hash shared across claims, snaks by property with their order
package dump

// Reference is one source of a claim: its hash (the same source on another
// claim has the same hash) and its snaks by property with the property order
// (snaks-order).
type Reference struct {
	Hash       string            `json:"hash"`
	Snaks      map[string][]Snak `json:"snaks"`
	SnaksOrder []string          `json:"snaks-order"`
}
