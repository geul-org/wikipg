//ff:type feature=load type=model
//ff:what table — one wikidata table the loader writes: its name and COPY columns
package load

// table is one table of sql/wikidata.sql the loader COPYs into, with the
// columns in the order rows give their values.
type table struct {
	Name    string
	Columns []string
}
