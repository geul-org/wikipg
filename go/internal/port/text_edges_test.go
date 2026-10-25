//ff:func feature=port type=encoder control=sequence
//ff:what TestTextEdges — one Edge per text with Edge Ids counting on from next
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestTextEdges(t *testing.T) {
	subj, _ := entityRef("Q1")
	next := uint64(4)
	rs := textEdges(edgeAlias, &next, subj, []source.Text{{Lang: "en", Value: "a"}, {Lang: "ko", Value: "b"}})
	if len(rs) != 2 || next != 6 || rs[1].Codec != edgeAlias || rs[1].Payload[1] != core.Word(5) {
		t.Fatalf("got %d next %d %+v", len(rs), next, rs)
	}
}
