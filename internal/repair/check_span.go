//ff:func feature=repair type=reader control=iteration dimension=1
//ff:what checkSpan — one batch end to end (read only): parse its lines, compare statement_reference and every reference occurrence with the database
package repair

import (
	"database/sql"
	"sort"
)

// checkSpan reads and parses sp (readSpan), compares its statement_reference
// rows with the stored ones (compareStatementRefs) and every reference
// occurrence with the stored rows of its hash (judge, values normalized by
// the database). It only reads.
func checkSpan(db *sql.DB, c Config, sp Span) (Found, error) {
	d, err := readSpan(db, c, sp)
	if err != nil {
		return Found{}, err
	}
	have, err := fetchStatementRefs(db, c, d.Lines)
	if err != nil {
		return Found{}, err
	}
	values, hashes := collectValues(d.Occs)
	stored, err := fetchStored(db, c, hashes)
	if err != nil {
		return Found{}, err
	}
	norm, err := normalizeValues(db, values)
	if err != nil {
		return Found{}, err
	}
	f := judge(d, stored, norm)
	f.Seq, f.Last, f.Entities, f.Bytes, f.Errors = sp.Seq, sp.Last, int64(len(d.Lines)), d.Bytes, d.Errors
	diff, lines := compareStatementRefs(d.Refs, have)
	f.StatementRefs, f.RefDiffs = int64(len(d.Refs)), diff
	for line := range lines {
		f.RefDiffEntities = append(f.RefDiffEntities, d.IDs[line])
	}
	sort.Strings(f.RefDiffEntities)
	return f, nil
}
