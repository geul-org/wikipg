//ff:type feature=dump type=model
//ff:what Entity — one Wikidata entity as the JSON dump writes it: ID, type, revision, terms, claims, sitelinks, and a lexeme's lemmas, language, category, forms and senses
package dump

// Entity is one line of the Wikidata JSON dump decoded (DecodeEntity): the
// entity ID and type (item, property), a property's datatype, the revision
// it was dumped at (lastrevid, modified), labels and descriptions by
// language, aliases by language in order, claims by property in order,
// sitelinks by site, and for a lexeme (lexeme dump) its lemmas by language
// variant, its language and lexical category (item IDs) and its forms and
// senses in order.
type Entity struct {
	ID           string                 `json:"id"`
	Type         string                 `json:"type"`
	Datatype     string                 `json:"datatype"`
	LastRevID    int64                  `json:"lastrevid"`
	Modified     string                 `json:"modified"`
	Labels       map[string]Term        `json:"labels"`
	Descriptions map[string]Term        `json:"descriptions"`
	Aliases      map[string][]Term      `json:"aliases"`
	Claims       map[string][]Statement `json:"claims"`
	Sitelinks    map[string]Sitelink    `json:"sitelinks"`

	Lemmas          map[string]Term `json:"lemmas"`
	Language        string          `json:"language"`
	LexicalCategory string          `json:"lexicalCategory"`
	Forms           []Form          `json:"forms"`
	Senses          []Sense         `json:"senses"`
}
