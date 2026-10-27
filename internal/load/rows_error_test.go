//ff:func feature=load type=model control=sequence
//ff:what TestRowsError — an error row keeps the first 1 KB without NUL bytes, NULL for an unknown entity, and counts
package load

import (
	"strings"
	"testing"
)

func TestRowsError(t *testing.T) {
	r := NewRows()
	r.Error(1, 7, "", "parse", "bad\x00", []byte(strings.Repeat("x", 2000)+"\x00"))
	r.Error(1, 8, "Q8", "value", "m", nil)
	row := r.byTable["load_error"][0]
	if r.Errors != 2 || row[2] != nil || row[4] != "bad" || len(row[5].(string)) != 1024 || r.byTable["load_error"][1][2] != "Q8" {
		t.Fatalf("got %+v", r.byTable["load_error"])
	}
}
