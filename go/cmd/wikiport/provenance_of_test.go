//ff:func feature=cli type=helper control=sequence
//ff:what TestProvenanceOf — the 20250908 dump, its acquisition time, the loader at 4225b3b, the porter and the library at the given commits
package main

import (
	"strings"
	"testing"

	"github.com/geul-org/wikiport/go/internal/port"
)

func TestProvenanceOf(t *testing.T) {
	p := provenanceOf("abc", "def")
	if !strings.Contains(p.DumpURI, "20250908") || p.AcquiredAt != 1758109791 || len(p.Tools) != 3 ||
		p.Tools[0] != (port.Tool{Doc: 3, Commit: "4225b3b4297796a79e95c4cda5aa54df3acc4069"}) || p.Tools[1] != (port.Tool{Doc: 1, Commit: "abc"}) || p.Tools[2] != (port.Tool{Doc: 2, Commit: "def"}) {
		t.Fatalf("got %+v", p)
	}
}
