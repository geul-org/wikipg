//ff:type feature=source type=model
//ff:what entityQueries — the SQL LoadEntity runs: existence, labels, aliases and claims of one entity
package source

// entityQueries is the SQL loadEntityWith runs for one entity ($1 = id, $2 =
// languages for labels and aliases). LoadEntity uses defaultQueries.
type entityQueries struct {
	Exists  string
	Labels  string
	Aliases string
	Claims  string
}
