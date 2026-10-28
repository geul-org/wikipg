//ff:func feature=load type=model control=sequence
//ff:what Rows.Table — the rows of one table in the batch, each its values in the table's column order
package load

// Table returns the rows of table in the batch in the order they were added,
// each its values in the table's column order (tables).
func (r *Rows) Table(table string) [][]any {
	return r.byTable[table]
}
