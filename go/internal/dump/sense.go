//ff:type feature=dump type=model
//ff:what Sense — one sense of a lexeme as the dump writes it: ID, glosses by language, claims
package dump

// Sense is one sense of a Wikidata lexeme (L4-S1): its glosses (short
// meanings) by language and its claims by property (P5137 item for this
// sense …).
type Sense struct {
	ID      string                 `json:"id"`
	Glosses map[string]Term        `json:"glosses"`
	Claims  map[string][]Statement `json:"claims"`
}
