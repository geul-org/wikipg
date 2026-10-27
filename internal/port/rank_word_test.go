//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestRankWord — normal 0, preferred 1, deprecated 2; other ranks fail
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRankWord(t *testing.T) {
	for rank, want := range map[string]core.Word{"normal": 0, "preferred": 1, "deprecated": 2} {
		if w, ok := rankWord(rank); !ok || w != want {
			t.Fatalf("%s: %d %v", rank, w, ok)
		}
	}
	if _, ok := rankWord(""); ok {
		t.Fatal("empty: want false")
	}
}
