//ff:func feature=repair type=command control=iteration dimension=1
//ff:what TestRunOrder — an order build: a dry run counts without writing; a write records each occurrence in another order, an entity range then a line range, and resumes past its checkpoint; reference_snak stays as it was and the permutations give each occurrence's dump order back
package repair

import (
	"context"
	"testing"
)

func TestRunOrder(t *testing.T) {
	db, c := loadOrderFixture(t)
	before := snakList(t, db, c, "hx") + "|" + snakList(t, db, c, "hy")
	st, err := Run(context.Background(), db, c)
	if err != nil || st.Spans != 2 || st.Entities != 3 || st.Orders != 2 || st.Written != 0 || st.Occurrences != (Tally{Same: 5, Order: 2}) || st.RefDiffs != 0 {
		t.Fatalf("dry run: %+v %v", st, err)
	}
	var n int
	if db.QueryRow("SELECT count(*) FROM information_schema.tables WHERE table_schema = $1 AND table_name LIKE 'reference_snak_order%'", c.Schema).Scan(&n); n != 0 {
		t.Fatal("dry run made tables")
	}
	c.DryRun, c.Kind, c.Lo, c.Hi = false, "Q", 3, 3
	if st, err = Run(context.Background(), db, c); err != nil || st.Written != 1 || st.Last != 4 {
		t.Fatalf("entities: %+v %v", st, err)
	}
	if got := orderList(t, db, c); got != "4.0/0={1,0}" {
		t.Fatalf("after Q3: %q", got)
	}
	c.Kind, c.Lo, c.Hi = "", 0, 0
	if st, err = Run(context.Background(), db, c); err != nil || st.Written != 2 || st.Last != 4 {
		t.Fatalf("lines: %+v %v", st, err)
	}
	if got := orderList(t, db, c); got != "3.1/0={2,3,1,0} 4.0/0={1,0}" {
		t.Fatalf("after all: %q", got)
	}
	if st, err = Run(context.Background(), db, c); err != nil || st.Planned != 0 || st.Last != 4 {
		t.Fatalf("resume: %+v %v", st, err)
	}
	if after := snakList(t, db, c, "hx") + "|" + snakList(t, db, c, "hy"); after != before {
		t.Fatalf("reference_snak changed: %q → %q", before, after)
	}
	want := map[int64]string{2: "P248 P813 P123 P123|P854 P248", 3: "P854|P123 P123 P813 P248", 4: "P248 P854"}
	for line, w := range want {
		var got string
		err := db.QueryRow("SELECT string_agg(props, '|' ORDER BY statement_id, ordinal) FROM (SELECT r.statement_id, r.ordinal, string_agg(s.property, ' ' ORDER BY coalesce(array_position(o.perm, s.ordinal) - 1, s.ordinal)) AS props "+
			"FROM "+c.Schema+".statement_reference r JOIN "+c.Schema+".reference_snak s ON s.hash = r.hash AND s.dump_id = 1 "+
			"LEFT JOIN "+c.Schema+".reference_snak_order o ON o.statement_id = r.statement_id AND o.ordinal = r.ordinal "+
			"WHERE r.statement_id >> 20 = (1::bigint << 35) | $1 AND s.property <> 'P143' GROUP BY r.statement_id, r.ordinal) x", line).Scan(&got)
		if err != nil || got != w {
			t.Fatalf("line %d in dump order: %q want %q (%v)", line, got, w, err)
		}
	}
}
