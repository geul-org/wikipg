//ff:type feature=load type=model
//ff:what Rows — the rows one batch writes, by table name
package load

// Rows collects the rows of one batch by table name, each row its values in
// the table's column order (tables). Errors counts the load_error rows.
type Rows struct {
	byTable map[string][][]any
	Errors  int
}
