//ff:func feature=repair type=helper control=iteration dimension=2
//ff:what permOf — the stored ordinal of each snak of an occurrence in its order: matches each to an unused stored row of the same content (ordinals aside)
package repair

// permOf returns, for each snak of want in order, the Ordinal of a row of
// have with the same content (Ordinal and ValueOrdinal aside), each row used
// once and the earliest unused one taken, so equal snaks keep their stored
// order. It is nil when some snak has no row left (not an Order verdict).
func permOf(want, have []Snak) []int {
	used := make([]bool, len(have))
	perm := make([]int, 0, len(want))
	for _, w := range want {
		w.Ordinal, w.ValueOrdinal = 0, 0
		j := -1
		for k, h := range have {
			h.Ordinal, h.ValueOrdinal = 0, 0
			if !used[k] && h == w {
				j = k
				break
			}
		}
		if j < 0 {
			return nil
		}
		used[j] = true
		perm = append(perm, have[j].Ordinal)
	}
	return perm
}
