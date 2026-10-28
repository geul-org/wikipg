//ff:func feature=repair type=helper control=iteration dimension=2
//ff:what collectValues — the distinct value texts (datavalue JSON) of the occurrences' snak rows, and their hashes
package repair

// collectValues returns the distinct non-NULL value texts (column value) of
// the rows of occs and the distinct hashes of occs, both in first-seen order.
func collectValues(occs []occ) ([]string, []string) {
	var values, hashes []string
	seenV, seenH := map[string]bool{}, map[string]bool{}
	for _, o := range occs {
		if !seenH[o.Hash] {
			seenH[o.Hash] = true
			hashes = append(hashes, o.Hash)
		}
		for _, row := range o.Rows {
			if v, ok := row[8].(string); ok && !seenV[v] {
				seenV[v] = true
				values = append(values, v)
			}
		}
	}
	return values, hashes
}
