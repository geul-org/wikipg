//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what Stats.add — adds one applied batch: its counts, the verdicts of the hashes processed, replacements and examples
package repair

import "fmt"

// add adds the applied batch f to st: lines, entities, bytes, occurrence
// verdicts and entities, statement_reference comparisons and errors (lists
// kept to 20), the verdict of each processed hash, the hashes replaced when
// written (a verdict other than Same) and up to 20 examples of differing
// hashes.
func (st *Stats) add(f Found, processed []HashCheck, skipped int64, written bool) {
	st.Spans, st.Last, st.Entities, st.Bytes = st.Spans+1, f.Last, st.Entities+f.Entities, st.Bytes+f.Bytes
	st.Occurrences.Same, st.Occurrences.Order = st.Occurrences.Same+f.Occurrences.Same, st.Occurrences.Order+f.Occurrences.Order
	st.Occurrences.Content, st.Occurrences.Missing = st.Occurrences.Content+f.Occurrences.Content, st.Occurrences.Missing+f.Occurrences.Missing
	st.ContentEntities, st.OrderEntities = st.ContentEntities+int64(len(f.ContentEntities)), st.OrderEntities+int64(len(f.OrderEntities))
	st.StatementRefs, st.RefDiffs, st.Skipped = st.StatementRefs+f.StatementRefs, st.RefDiffs+f.RefDiffs, st.Skipped+skipped
	st.RefDiffEntities = append(st.RefDiffEntities, f.RefDiffEntities[:min(len(f.RefDiffEntities), max(20-len(st.RefDiffEntities), 0))]...)
	st.Errors = append(st.Errors, f.Errors[:min(len(f.Errors), max(20-len(st.Errors), 0))]...)
	for _, h := range processed {
		st.Hashes.add(h.Verdict)
		if h.Verdict != Same && written {
			st.Replaced++
		}
		if h.Verdict != Same && len(st.Examples) < 20 {
			st.Examples = append(st.Examples, fmt.Sprintf("%s %s %s (statement %d reference %d)", h.Hash, h.Entity, h.Verdict, h.At.Statement, h.At.Ordinal))
		}
	}
}
