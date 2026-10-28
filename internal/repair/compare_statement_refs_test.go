//ff:func feature=repair type=helper control=sequence
//ff:what TestCompareStatementRefs — a row only in the dump, one only stored and one with another hash count, with their lines
package repair

import "testing"

func TestCompareStatementRefs(t *testing.T) {
	k := func(line, i int64) int64 { return 1<<55 | line<<20 | i }
	want := []stmtRef{{k(2, 0), 0, "a"}, {k(2, 0), 1, "b"}, {k(3, 1), 0, "c"}, {k(5, 0), 0, "e"}}
	have := []stmtRef{{k(2, 0), 0, "a"}, {k(2, 0), 1, "x"}, {k(4, 0), 0, "d"}, {k(5, 0), 0, "e"}}
	diff, lines := compareStatementRefs(want, have)
	if diff != 3 || len(lines) != 3 || !lines[2] || !lines[3] || !lines[4] || lines[5] {
		t.Fatalf("%d %v", diff, lines)
	}
}
