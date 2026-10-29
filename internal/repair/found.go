//ff:type feature=repair type=model
//ff:what Found — a checked batch: its place, entities and bytes, each hash's first occurrence checked, occurrence and statement_reference comparisons, occurrences in another order, errors
package repair

// Found is the result of checking one batch (checkSpan): its Seq and Last
// line, the entities and raw bytes read, the first occurrence of each hash
// with its verdict (Hashes, in dump order), every occurrence compared with
// the rows stored when the batch was read (Occurrences) with the entities
// that have an occurrence differing in content or only in order, the
// occurrences differing only in order with their permutation (Orders),
// statement_reference rows compared with the dump (StatementRefs) with the
// rows that differ and their entities, and lines that did not parse.
type Found struct {
	Seq             int64
	Last            int64
	Entities        int64
	Bytes           int64
	Hashes          []HashCheck
	Occurrences     Tally
	ContentEntities []string
	OrderEntities   []string
	Orders          []OccOrder
	StatementRefs   int64
	RefDiffs        int64
	RefDiffEntities []string
	Errors          []string
}
