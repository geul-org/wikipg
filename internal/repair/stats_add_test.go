//ff:func feature=repair type=helper control=sequence
//ff:what TestStatsAdd — batch counts add up, lists stop at 20, replacements count only when written, examples name differing hashes
package repair

import (
	"strings"
	"testing"
)

func TestStatsAdd(t *testing.T) {
	var st Stats
	many := make([]string, 25)
	f := Found{Last: 9, Entities: 2, Bytes: 10, Occurrences: Tally{Same: 1, Order: 2, Content: 3, Missing: 4}, ContentEntities: []string{"Q1"}, OrderEntities: []string{"Q1", "Q2"},
		StatementRefs: 5, RefDiffs: 1, RefDiffEntities: many, Errors: many}
	checks := []HashCheck{{Hash: "s"}, {Hash: "o", Entity: "Q2", Verdict: Order}}
	st.add(f, checks, 1, false)
	st.add(f, checks, 2, true)
	if st.Spans != 2 || st.Last != 9 || st.Entities != 4 || st.Bytes != 20 || st.Occurrences != (Tally{Same: 2, Order: 4, Content: 6, Missing: 8}) || st.ContentEntities != 2 || st.OrderEntities != 4 {
		t.Fatalf("%+v", st)
	}
	if st.StatementRefs != 10 || st.RefDiffs != 2 || len(st.RefDiffEntities) != 20 || len(st.Errors) != 20 || st.Skipped != 3 || st.Hashes != (Tally{Same: 2, Order: 2}) || st.Replaced != 1 {
		t.Fatalf("%+v", st)
	}
	if len(st.Examples) != 2 || !strings.HasPrefix(st.Examples[0], "o Q2 order") {
		t.Fatalf("examples %v", st.Examples)
	}
	full := Stats{Examples: many[:20]}
	full.add(Found{}, checks, 0, false)
	if len(full.Examples) != 20 {
		t.Fatal("examples bound")
	}
}
