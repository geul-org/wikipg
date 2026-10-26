//ff:type feature=dump type=model
//ff:what Form — one form of a lexeme as the dump writes it: ID, representations by language, grammatical features, claims
package dump

// Form is one form of a Wikidata lexeme (L4-F1): its spellings by language
// variant, its grammatical features (item IDs such as Q146786 plural) and its
// claims by property.
type Form struct {
	ID                  string                 `json:"id"`
	Representations     map[string]Term        `json:"representations"`
	GrammaticalFeatures []string               `json:"grammaticalFeatures"`
	Claims              map[string][]Statement `json:"claims"`
}
