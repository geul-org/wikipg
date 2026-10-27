//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddText — text goes last in the row; NUL bytes are removed with an error
package load

import "testing"

func TestRowsAddText(t *testing.T) {
	r := NewRows()
	r.addText(Config{DumpID: 1}, 3, "Q1", "label", "a\x00b", "Q1", int16(1), "en")
	row := r.byTable["label"][0]
	if len(row) != 4 || row[3] != "ab" || r.Errors != 1 {
		t.Fatalf("got %v %d", row, r.Errors)
	}
}
