//ff:func feature=repair type=helper control=sequence
//ff:what TestColumnsOf — rows of several checks in order, column by column, NULLs kept
package repair

import "testing"

func TestColumnsOf(t *testing.T) {
	row := func(h string, ord int, v any) []any {
		return []any{h, int16(1), ord, "P1", 0, "value", "string", "string", v, nil, nil}
	}
	c := columnsOf([]HashCheck{{Rows: [][]any{row("a", 0, `"x"`), row("a", 1, nil)}}, {}, {Rows: [][]any{row("b", 0, `"y"`)}}})
	if len(c.Hash) != 3 || c.Hash[2] != "b" || c.Ordinal[1] != 1 || c.DumpID[0] != 1 || c.Value[1].Valid || c.Value[2].String != `"y"` || c.ValueEntity[0].Valid || c.SnakType[0] != "value" {
		t.Fatalf("%+v", c)
	}
}
