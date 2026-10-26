//ff:func feature=load type=model control=sequence
//ff:what NewRows — an empty batch of rows
package load

// NewRows returns an empty Rows for one batch.
func NewRows() *Rows {
	return &Rows{byTable: map[string][][]any{}}
}
