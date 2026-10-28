//ff:func feature=repair type=helper control=sequence
//ff:what TestCollectValues — distinct values and hashes in first-seen order, NULL values left out
package repair

import (
	"slices"
	"testing"
)

func TestCollectValues(t *testing.T) {
	row := func(v any) []any { return []any{"", int16(1), 0, "P", 0, "value", nil, nil, v, nil, nil} }
	occs := []occ{{Hash: "a", Rows: [][]any{row("1"), row(nil)}}, {Hash: "b", Rows: [][]any{row("2"), row("1")}}, {Hash: "a", Rows: [][]any{row("3")}}}
	values, hashes := collectValues(occs)
	if !slices.Equal(values, []string{"1", "2", "3"}) || !slices.Equal(hashes, []string{"a", "b"}) {
		t.Fatalf("%v %v", values, hashes)
	}
}
