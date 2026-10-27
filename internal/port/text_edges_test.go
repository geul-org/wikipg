//ff:func feature=port type=encoder control=sequence
//ff:what TestTextEdges — one Edge per text with Edge Ids from the assignment; a malformed language is skipped and counted
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikiport/internal/source"
)

func TestTextEdges(t *testing.T) {
	subj, _ := entityRef("Q1")
	ids := freshIDs(4)
	rep := Report{Skipped: map[string]int{}}
	rs := textEdges(edgeAlias, ids, subj, []source.Text{{Lang: "en", Value: "a"}, {Lang: "x y", Value: "c"}, {Lang: "ko", Value: "b"}}, &rep)
	if len(rs) != 2 || ids.Next() != 6 || rs[1].Codec != edgeAlias || rs[1].Payload[1] != core.Word(5) || rep.Skipped["언어 코드: x y"] != 1 {
		t.Fatalf("got %d next %d %+v %v", len(rs), ids.Next(), rs, rep.Skipped)
	}
}
