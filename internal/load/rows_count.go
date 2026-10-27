//ff:func feature=load type=model control=sequence
//ff:what Rows.Count — how many rows a table has in the batch
package load

// Count returns the number of rows of table in the batch.
func (r *Rows) Count(table string) int {
	return len(r.byTable[table])
}
