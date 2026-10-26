//ff:func feature=load type=model control=sequence
//ff:what Rows.add — appends one row to a table
package load

// add appends one row, its values in the table's column order, to table.
func (r *Rows) add(table string, values ...any) {
	r.byTable[table] = append(r.byTable[table], values)
}
