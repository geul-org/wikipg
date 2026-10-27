//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddStatements — statements in sorted property order with keys from the line and start, returning the next number; an out-of-range key is one error
package load

import "testing"

func TestRowsAddStatements(t *testing.T) {
	r := NewRows()
	e := sampleEntity()
	next, ok := r.addStatements(Config{DumpID: 1}, 9, e.ID, e.Claims, 0, NewRefSet(64))
	st := r.byTable["statement"]
	if !ok || next != 3 || len(st) != 3 || st[0][3] != "P2048" || st[1][3] != "P40" || st[2][3] != "P69" || st[0][0] != int64(1<<55|9<<20) || st[2][0] != int64(1<<55|9<<20|2) {
		t.Fatalf("got %v %d %v", st, next, ok)
	}
	if next, ok := r.addStatements(Config{DumpID: 1}, 9, "L1-F1", e.Claims, 5, NewRefSet(64)); !ok || next != 8 || r.byTable["statement"][3][0] != int64(1<<55|9<<20|5) || r.byTable["statement"][3][2] != "L1-F1" {
		t.Fatalf("from 5: %d %v", next, ok)
	}
	bad := NewRows()
	if _, ok := bad.addStatements(Config{DumpID: 0}, 9, e.ID, e.Claims, 0, NewRefSet(64)); ok || bad.Count("statement") != 0 || bad.Errors != 1 {
		t.Fatalf("bad %+v", bad.byTable)
	}
}
