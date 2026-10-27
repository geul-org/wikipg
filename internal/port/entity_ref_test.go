//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestEntityRef — Q and P IDs become WIKIDATA-QID and WIKIDATA-PID Records; lexemes and bad IDs do not
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEntityRef(t *testing.T) {
	q, ok := entityRef("Q42")
	if ws, _ := core.EncodeRecord(q); !ok || !slices.Equal(ws, []core.Word{0x00A0, 42}) {
		t.Fatalf("Q42: %04X %v", ws, ok)
	}
	p, ok := entityRef("P70000")
	if ws, _ := core.EncodeRecord(p); !ok || !slices.Equal(ws, []core.Word{0x0128, 1, 0x1170}) {
		t.Fatalf("P70000: %04X %v", ws, ok)
	}
	for _, id := range []string{"L5", "E10", "Q0", ""} {
		if _, ok := entityRef(id); ok {
			t.Fatalf("%q: want false", id)
		}
	}
}
