//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddStatement — the statement row keeps GUID, rank and main snak; qualifiers follow qualifiers-order
package load

import "testing"

func TestRowsAddStatement(t *testing.T) {
	r := NewRows()
	s := sampleEntity().Claims["P69"][0]
	r.addStatement(Config{DumpID: 1}, 9, "Q42", 77, 0, s, NewRefSet(64))
	st := r.byTable["statement"][0]
	if st[0] != int64(77) || st[1] != "Q42$A" || st[5] != "normal" || st[10] != "Q691283" || st[11] != int16(1) {
		t.Fatalf("statement %v", st)
	}
	q := r.byTable["qualifier"]
	if len(q) != 2 || q[0][2] != "P580" || q[0][1] != 0 || q[1][2] != "P582" || q[0][9] != "hP580" {
		t.Fatalf("qualifiers %v", q)
	}
}
