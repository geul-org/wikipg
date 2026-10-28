//ff:type feature=repair type=model
//ff:what Tally — counts by verdict
package repair

// Tally counts comparisons by verdict.
type Tally struct {
	Same    int64
	Order   int64
	Content int64
	Missing int64
}
