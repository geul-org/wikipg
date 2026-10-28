//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what compareSnaks — how stored rows compare with an occurrence's list: same, order only (same multiset), content, or missing
package repair

import "slices"

// compareSnaks compares the stored rows have of a hash with the list want of
// one occurrence: Same when they match row for row, Missing when nothing is
// stored for a list with snaks, Order when both hold the same snaks (ordinals
// aside) as a multiset and Content otherwise.
func compareSnaks(want, have []Snak) Verdict {
	if slices.Equal(want, have) {
		return Same
	}
	if len(have) == 0 {
		return Missing
	}
	count := map[Snak]int{}
	for _, s := range want {
		s.Ordinal, s.ValueOrdinal = 0, 0
		count[s]++
	}
	for _, s := range have {
		s.Ordinal, s.ValueOrdinal = 0, 0
		count[s]--
		if count[s] < 0 {
			return Content
		}
	}
	if len(want) != len(have) {
		return Content
	}
	return Order
}
