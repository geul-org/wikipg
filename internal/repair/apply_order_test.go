//ff:func feature=repair type=writer control=iteration dimension=1
//ff:what TestApplyOrder — a dry run counts without writing; a write stores the orders and the checkpoint in one transaction; each failing step leaves nothing behind
package repair

import "testing"

func TestApplyOrder(t *testing.T) {
	db, c := loadOrderFixture(t)
	f, err := checkSpan(db, c, Span{From: 2, To: 4, Last: 4})
	if err != nil {
		t.Fatal(err)
	}
	var st Stats
	if err := applyOrder(db, c, f, &st); err != nil || st.Orders != 2 || st.Written != 0 || st.Spans != 1 {
		t.Fatalf("dry: %+v %v", st, err)
	}
	c.DryRun = false
	if err := applyOrder(db, c, f, &st); err == nil {
		t.Fatal("no tables: want error")
	}
	if err := ensureTables(db, c); err != nil {
		t.Fatal(err)
	}
	if err := applyOrder(db, c, f, &st); err != nil || st.Orders != 4 || st.Written != 2 {
		t.Fatalf("write: %+v %v", st, err)
	}
	if got := orderList(t, db, c); got != "3.1/0={2,3,1,0} 4.0/0={1,0}" {
		t.Fatalf("stored %q", got)
	}
	if n, _ := readCheckpoint(db, c); n != 4 {
		t.Fatalf("checkpoint %d", n)
	}
	s := c.Schema + "."
	steps := []string{
		"CREATE FUNCTION " + s + "refuse() RETURNS trigger LANGUAGE plpgsql AS $$BEGIN RAISE EXCEPTION 'refused at commit'; END$$; " +
			"CREATE CONSTRAINT TRIGGER refuse AFTER INSERT OR UPDATE ON " + s + "reference_snak_order DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION " + s + "refuse()",
		"ALTER TABLE " + s + "reference_snak_order_checkpoint RENAME TO cp_gone",
		"ALTER TABLE " + s + "reference_snak_order RENAME TO order_gone",
	}
	f.Orders[0].Perm, f.Last = []int{3, 2, 1, 0}, 9
	for _, q := range steps {
		if _, err := db.Exec(q); err != nil {
			t.Fatal(q, err)
		}
		if err := applyOrder(db, c, f, &st); err == nil {
			t.Fatalf("after %s: want error", q)
		}
	}
	var perm string
	db.QueryRow("SELECT perm::text FROM " + s + "order_gone ORDER BY statement_id LIMIT 1").Scan(&perm)
	if perm != "{2,3,1,0}" || st.Written != 2 {
		t.Fatalf("a failed batch wrote: %s %+v", perm, st)
	}
	if err := applyOrder(closedDB(t), c, f, &st); err == nil {
		t.Fatal("closed: want error")
	}
}
