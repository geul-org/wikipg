//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddStatements — statements in sorted property order with keys from the line; an out-of-range key is one error
package load

import "testing"

func TestRowsAddStatements(t *testing.T) {
	r := NewRows()
	r.addStatements(Config{DumpID: 1}, 9, sampleEntity(), NewRefSet(64))
	st := r.byTable["statement"]
	if len(st) != 3 || st[0][3] != "P2048" || st[1][3] != "P40" || st[2][3] != "P69" || st[0][0] != int64(1<<55|9<<20) || st[2][0] != int64(1<<55|9<<20|2) {
		t.Fatalf("got %v", st)
	}
	bad := NewRows()
	bad.addStatements(Config{DumpID: 0}, 9, sampleEntity(), NewRefSet(64))
	if bad.Count("statement") != 0 || bad.Errors != 1 {
		t.Fatalf("bad %+v", bad.byTable)
	}
}
