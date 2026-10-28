//ff:type feature=repair type=model
//ff:what Verdict — how the stored snak list of a hash compares with an occurrence's: same, order only, content, missing
package repair

// Verdict is how the reference_snak rows of a hash compare with the snak list
// of one occurrence in the dump.
type Verdict int

// Verdicts: Same (row for row), Order (the same snaks in another order),
// Content (other snaks: some missing, some extra or doubled), Missing (no
// rows stored for a source that has snaks).
const (
	Same Verdict = iota
	Order
	Content
	Missing
)
