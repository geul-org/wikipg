//ff:func feature=port type=encoder control=sequence
//ff:what TestTextEdges — one Edge per text with Edge Ids counting on from next; a malformed language is skipped and counted
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/wikiport/go/internal/source"
)

func TestTextEdges(t *testing.T) {
	subj, _ := entityRef("Q1")
	next := uint64(4)
	rep := Report{Skipped: map[string]int{}}
	rs := textEdges(edgeAlias, &next, subj, []source.Text{{Lang: "en", Value: "a"}, {Lang: "x y", Value: "c"}, {Lang: "ko", Value: "b"}}, &rep)
	if len(rs) != 2 || next != 6 || rs[1].Codec != edgeAlias || rs[1].Payload[1] != core.Word(5) || rep.Skipped["언어 코드: x y"] != 1 {
		t.Fatalf("got %d next %d %+v %v", len(rs), next, rs, rep.Skipped)
	}
}
