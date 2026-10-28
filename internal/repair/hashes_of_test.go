//ff:func feature=repair type=helper control=sequence
//ff:what TestHashesOf — the hashes in order; none gives an empty list
package repair

import (
	"slices"
	"testing"
)

func TestHashesOf(t *testing.T) {
	if got := hashesOf([]HashCheck{{Hash: "b"}, {Hash: "a"}}); !slices.Equal(got, []string{"b", "a"}) {
		t.Fatal(got)
	}
	if got := hashesOf(nil); got == nil || len(got) != 0 {
		t.Fatal("empty")
	}
}
