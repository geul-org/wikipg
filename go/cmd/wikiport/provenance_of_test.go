//ff:func feature=cli type=helper control=sequence
//ff:what TestProvenanceOf — the 20250908 dump, its acquisition time, the loader at 4225b3b and the porter at the given commit
package main

import (
	"strings"
	"testing"
)

func TestProvenanceOf(t *testing.T) {
	p := provenanceOf("abc")
	if !strings.Contains(p.DumpURI, "20250908") || p.AcquiredAt != 1758109791 || !strings.HasPrefix(p.Loader.Version, "4225b3b") || p.Porter.Version != "abc" {
		t.Fatalf("got %+v", p)
	}
}
