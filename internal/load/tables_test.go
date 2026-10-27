//ff:func feature=load type=model control=iteration dimension=1
//ff:what TestTables — twelve tables, unique names, entity_raw first and load_error last, snak tables share the snak columns
package load

import (
	"slices"
	"testing"
)

func TestTables(t *testing.T) {
	ts := tables()
	seen := map[string]bool{}
	for _, x := range ts {
		if seen[x.Name] || len(x.Columns) == 0 {
			t.Fatalf("table %+v", x)
		}
		seen[x.Name] = true
	}
	if len(ts) != 18 || ts[0].Name != "entity_raw" || ts[11].Name != "lexeme" || ts[17].Name != "load_error" {
		t.Fatalf("got %d tables", len(ts))
	}
	for _, i := range []int{6, 7, 10} {
		if !slices.Contains(ts[i].Columns, "value_entity") || !slices.Contains(ts[i].Columns, "snaktype") {
			t.Fatalf("%s lacks snak columns", ts[i].Name)
		}
	}
}
