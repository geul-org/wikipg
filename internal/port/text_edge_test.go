//ff:func feature=port type=encoder control=sequence
//ff:what TestTextEdge — Edge Id, subject, then the language text (the edge/ALIAS example words); a malformed language is false
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikiport/internal/source"
)

func TestTextEdge(t *testing.T) {
	subj, _ := entityRef("Q42")
	r, ok := textEdge(edgeAlias, freshIDs(3), subj, source.Text{Lang: "en", Value: "DNA"})
	ws, _ := core.EncodeRecord(r)
	if !ok || !slices.Equal(ws, []core.Word{0x1485, 0x0088, 0x0003, 0x00A0, 0x002A, 0x0A4B, 0x08C6, 0x656E, 0x444E, 0x4100}) {
		t.Fatalf("got %04X", ws)
	}
	if _, ok := textEdge(edgeLabel, freshIDs(1), subj, source.Text{Lang: "x y", Value: "z"}); ok {
		t.Fatal("bad language: want false")
	}
}
