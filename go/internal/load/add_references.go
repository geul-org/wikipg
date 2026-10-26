//ff:func feature=load type=model control=iteration dimension=2
//ff:what Rows.addReferences — a statement's reference links in order, and each source's snaks the first time its hash is seen
package load

import "github.com/geul-org/wikiport/go/internal/dump"

// addReferences links statement sid to each of its references in order
// (statement_reference). A source whose hash refs has not seen gets its
// reference row and its snaks in dump order (reference_snak); one seen before
// is only linked.
func (r *Rows) addReferences(c Config, line int64, id string, sid int64, rs []dump.Reference, refs *RefSet) {
	for ord, ref := range rs {
		r.add("statement_reference", sid, ord, ref.Hash)
		if !refs.First(ref.Hash) {
			continue
		}
		r.add("reference", ref.Hash, c.DumpID)
		for _, sn := range orderedSnaks(ref.SnaksOrder, ref.Snaks) {
			vals := r.snakValues(c, line, id, sn.Snak)
			r.add("reference_snak", append([]any{ref.Hash, c.DumpID, sn.Ordinal, sn.Snak.Property, sn.ValueOrdinal}, vals...)...)
		}
	}
}
