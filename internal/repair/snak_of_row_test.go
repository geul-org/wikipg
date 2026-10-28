//ff:func feature=repair type=helper control=sequence
//ff:what TestSnakOfRow — loader row columns map to Snak fields, NULL to invalid, the value to its normalized form
package repair

import (
	"database/sql"
	"testing"
)

func TestSnakOfRow(t *testing.T) {
	row := []any{"h", int16(1), 2, "P248", 1, "value", "wikibase-item", "wikibase-entityid", `{"id": "Q1"}`, "Q1", nil}
	s := snakOfRow(row, map[string]string{`{"id": "Q1"}`: `{"id":"Q1"}`})
	want := Snak{Ordinal: 2, ValueOrdinal: 1, Property: "P248", SnakType: "value", Datatype: sql.NullString{String: "wikibase-item", Valid: true},
		ValueType: sql.NullString{String: "wikibase-entityid", Valid: true}, Value: sql.NullString{String: `{"id":"Q1"}`, Valid: true}, ValueEntity: sql.NullString{String: "Q1", Valid: true}}
	if s != want {
		t.Fatalf("%+v", s)
	}
	if s := snakOfRow([]any{"h", int16(1), 0, "P1", 0, "novalue", nil, nil, nil, nil, nil}, nil); s.Value.Valid || s.Datatype.Valid {
		t.Fatalf("NULLs: %+v", s)
	}
}
