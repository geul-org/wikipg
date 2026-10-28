//ff:func feature=repair type=helper control=selection
//ff:what Tally.add — counts one verdict
package repair

// add counts one comparison with verdict v.
func (t *Tally) add(v Verdict) {
	switch v {
	case Same:
		t.Same++
	case Order:
		t.Order++
	case Content:
		t.Content++
	default:
		t.Missing++
	}
}
