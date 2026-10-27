//ff:func feature=load type=helper control=sequence
//ff:what TestNulJSON — a string with NUL becomes its JSON literal; any other string is NULL
package load

import "testing"

func TestNulJSON(t *testing.T) {
	if got := nulJSON("a\x00b"); got != `"a\u0000b"` {
		t.Fatalf("nul: %v", got)
	}
	if got := nulJSON("plain"); got != nil {
		t.Fatalf("plain: %v", got)
	}
}
