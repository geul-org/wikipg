//ff:func feature=repair type=reader control=sequence
//ff:what TestNormalizeValues — dump JSON maps to jsonb text form (key order, spacing, escapes); none gives an empty map; invalid JSON fails
package repair

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestNormalizeValues(t *testing.T) {
	db, _ := testSchema(t)
	got, err := normalizeValues(db, []string{`{"id":"Q1","entity-type":"item"}`, `"é"`})
	if err != nil || got[`{"id":"Q1","entity-type":"item"}`] != `{"id": "Q1", "entity-type": "item"}` || got[`"é"`] != `"é"` {
		t.Fatalf("%v %v", got, err)
	}
	if got, err := normalizeValues(db, nil); err != nil || len(got) != 0 {
		t.Fatalf("none: %v %v", got, err)
	}
	if _, err := normalizeValues(db, []string{"{bad"}); err == nil {
		t.Fatal("invalid JSON: want error")
	}
}
