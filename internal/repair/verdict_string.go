//ff:func feature=repair type=helper control=selection
//ff:what Verdict.String — the verdict's name as stored in repair_reference_snak_done (same, order, content, missing)
package repair

// String names v as repair_reference_snak_done.verdict does.
func (v Verdict) String() string {
	switch v {
	case Same:
		return "same"
	case Order:
		return "order"
	case Content:
		return "content"
	default:
		return "missing"
	}
}
