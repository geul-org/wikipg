//ff:type feature=source type=model
//ff:what Text — one language-tagged string of an entity: a label or an alias
package source

// Text is one language-tagged string of a Wikidata entity (a label or an
// alias). Lang is the Wikidata language code (ko, en).
type Text struct {
	Lang  string
	Value string
}
