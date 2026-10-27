//ff:type feature=source type=model
//ff:what entityQueries — the SQL LoadEntity runs: existence, labels, descriptions, aliases, claims and qualifiers of one entity
package source

// entityQueries is the SQL loadEntityWith runs for one entity ($1 = id, $2 =
// languages for labels, descriptions and aliases). LoadEntity uses queriesFor.
type entityQueries struct {
	Exists       string
	Labels       string
	Descriptions string
	Aliases      string
	Claims       string
	Qualifiers   string
}
