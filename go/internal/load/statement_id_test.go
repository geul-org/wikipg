//ff:func feature=load type=helper control=iteration dimension=1
//ff:what TestStatementID — dump, line and index pack into one key; out-of-range dump, line or index fail
package load

import "testing"

func TestStatementID(t *testing.T) {
	if k, err := statementID(1, 2, 3); err != nil || k != 1<<55|2<<20|3 {
		t.Fatalf("%d %v", k, err)
	}
	if k, err := statementID(255, 1<<35-1, 1<<20-1); err != nil || k < 0 {
		t.Fatalf("max: %d %v", k, err)
	}
	for _, c := range [][3]int64{{0, 1, 0}, {256, 1, 0}, {1, -1, 0}, {1, 1 << 35, 0}, {1, 1, -1}, {1, 1, 1 << 20}} {
		if _, err := statementID(int16(c[0]), c[1], int(c[2])); err == nil {
			t.Fatalf("%v: want error", c)
		}
	}
}
