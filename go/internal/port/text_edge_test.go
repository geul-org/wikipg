//ff:func feature=port type=encoder control=sequence
//ff:what TestTextEdge — a LABEL is Edge Id, LANGUAGE-BCP47-CODE, Subject and UTF8-NFC-TEXT
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestTextEdge(t *testing.T) {
	subj, _ := entityRef("Q42")
	r := textEdge(edgeLabel, 1, subj, source.Text{Lang: "en", Value: "Douglas Adams"})
	want := []core.Word{0x0088, 1, 0x08C6, 0x656E, 0x00A0, 42, 0x0BC3, 0x446F, 0x7567, 0x6C61, 0x7320, 0x4164, 0x616D, 0x7300}
	if r.Kind != core.KindEdge || r.Codec != edgeLabel || !slices.Equal(r.Payload, want) {
		t.Fatalf("got %04X", r.Payload)
	}
}
