//ff:type feature=source type=model
//ff:what Qualifier — one Wikidata qualifier snak of a loaded dump: statement, property, value type, value
package source

// Qualifier is one qualifier snak of a loaded dump (wiki/wikidata.sql
// qualifier): the statement_id it qualifies, its property, the datavalue type
// and the value in the same form as a Claim's (an entity is a bare ID; empty
// for novalue and somevalue).
type Qualifier struct {
	StatementID int64
	Property    string
	Type        string
	Value       string
}
