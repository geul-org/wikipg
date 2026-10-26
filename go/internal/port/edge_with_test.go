//ff:func feature=port type=encoder control=sequence
//ff:what TestEdgeWith — the Edge Id comes from the assignment by content: equal content in the previous version keeps its ID, records then raw tail words follow it
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestEdgeWith(t *testing.T) {
	subj, _ := entityRef("Q42")
	r, id := edgeWith(freshIDs(7), edgeTriple, []core.Record{subj}, 1)
	if id != 7 || !slices.Equal(r.Payload, []core.Word{0x0088, 7, 0x00A0, 42, 1}) {
		t.Fatalf("got %d %04X", id, r.Payload)
	}
}
