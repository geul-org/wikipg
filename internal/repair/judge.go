//ff:func feature=repair type=helper control=iteration dimension=2
//ff:what judge — compares every occurrence of a batch with the stored rows of its hash; the first occurrence of each hash becomes its HashCheck
package repair

// judge compares each occurrence of d with the stored rows of its hash
// (stored, values in jsonb text form by norm): every occurrence is counted in
// Occurrences (its entity listed once per kind of difference), and the first
// occurrence of each hash in the batch (d.Occs is in dump order) becomes a
// HashCheck with its rows.
func judge(d spanData, stored map[string][]Snak, norm map[string]string) Found {
	f := Found{}
	seen := map[string]bool{}
	content, order := map[string]bool{}, map[string]bool{}
	for _, o := range d.Occs {
		var want []Snak
		for _, row := range o.Rows {
			want = append(want, snakOfRow(row, norm))
		}
		v := compareSnaks(want, stored[o.Hash])
		f.Occurrences.add(v)
		if v == Content && !content[o.Entity] {
			content[o.Entity] = true
			f.ContentEntities = append(f.ContentEntities, o.Entity)
		}
		if v == Order && !order[o.Entity] {
			order[o.Entity] = true
			f.OrderEntities = append(f.OrderEntities, o.Entity)
		}
		if !seen[o.Hash] {
			seen[o.Hash] = true
			f.Hashes = append(f.Hashes, HashCheck{Hash: o.Hash, At: o.At, Entity: o.Entity, Verdict: v, Rows: o.Rows})
		}
	}
	return f
}
