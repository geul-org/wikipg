//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddText — text and its JSON go last in the row; NUL bytes leave the text but stay in value_json; plain text has no JSON
package load

import "testing"

func TestRowsAddText(t *testing.T) {
	r := NewRows()
	r.addText("label", "a\x00b", "Q1", int16(1), "en")
	r.addText("label", "plain", "Q2", int16(1), "en")
	nul, plain := r.byTable["label"][0], r.byTable["label"][1]
	if len(nul) != 5 || nul[3] != "ab" || nul[4] != `"a\u0000b"` || r.Errors != 0 {
		t.Fatalf("nul: %v %d", nul, r.Errors)
	}
	if plain[3] != "plain" || plain[4] != nil {
		t.Fatalf("plain: %v", plain)
	}
}
