//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddReferences — every reference is linked; a source's snaks are written with their occurrence the first time, and again only for an earlier occurrence
package load

import "testing"

func TestRowsAddReferences(t *testing.T) {
	r := NewRows()
	refs := NewRefSet(64)
	rs := sampleEntity().Claims["P69"][0].References
	r.addReferences(Config{DumpID: 1}, 9, "Q42", 5, rs, refs)
	r.addReferences(Config{DumpID: 1}, 9, "Q42", 7, rs, refs)
	if r.Count("statement_reference") != 2 || r.Count("reference") != 1 || r.Count("reference_snak") != 1 {
		t.Fatalf("got %+v", r.byTable)
	}
	if s := r.byTable["reference_snak"][0]; len(s) != 13 || s[0] != "r1" || s[3] != "P248" || s[9] != "Q5375741" || s[11] != int64(5) || s[12] != 0 {
		t.Fatalf("snak %v", s)
	}
	r.addReferences(Config{DumpID: 1}, 9, "Q42", 3, rs, refs)
	if r.Count("reference_snak") != 2 || r.byTable["reference_snak"][1][11] != int64(3) {
		t.Fatalf("earlier occurrence: %v", r.byTable["reference_snak"])
	}
}
