//ff:func feature=load type=helper control=sequence
//ff:what TestRowsSnakValues — a value snak gives type, JSON and entity; novalue gives NULLs; a \u0000 value is NULL with an error
package load

import (
	"encoding/json"
	"testing"

	"github.com/geul-org/wikiport/go/internal/dump"
)

func TestRowsSnakValues(t *testing.T) {
	r := NewRows()
	c := Config{DumpID: 1}
	v := r.snakValues(c, 2, "Q1", dump.Snak{SnakType: "value", Property: "P31", Datatype: "wikibase-item", DataValue: &dump.DataValue{Type: "wikibase-entityid", Value: json.RawMessage(`{"id":"Q5"}`)}})
	if v[0] != "value" || v[1] != "wikibase-item" || v[2] != "wikibase-entityid" || v[3] != `{"id":"Q5"}` || v[4] != "Q5" {
		t.Fatalf("value: %v", v)
	}
	n := r.snakValues(c, 2, "Q1", dump.Snak{SnakType: "novalue", Property: "P40"})
	if n[0] != "novalue" || n[1] != nil || n[2] != nil || n[3] != nil || n[4] != nil {
		t.Fatalf("novalue: %v", n)
	}
	z := r.snakValues(c, 2, "Q1", dump.Snak{SnakType: "value", Property: "P1", DataValue: &dump.DataValue{Type: "string", Value: json.RawMessage(`"a\u0000"`)}})
	if z[2] != "string" || z[3] != nil || r.Errors != 1 {
		t.Fatalf("nul: %v %d", z, r.Errors)
	}
}
