//ff:func feature=port type=encoder control=sequence
//ff:what TestTripleEdge — Edge Id, Predicate, Subject, Object, then the rank word
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestTripleEdge(t *testing.T) {
	pred, _ := entityRef("P31")
	subj, _ := entityRef("Q42")
	obj, _ := entityRef("Q5")
	r := tripleEdge(3, pred, subj, obj, 1)
	if r.Codec != edgeTriple || !slices.Equal(r.Payload, []core.Word{0x0088, 3, 0x00A8, 31, 0x00A0, 42, 0x00A0, 5, 1}) {
		t.Fatalf("got %04X", r.Payload)
	}
}
