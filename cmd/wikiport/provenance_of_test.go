//ff:func feature=cli type=helper control=sequence
//ff:what TestProvenanceOf — the dump's URL and download time, wikiload, the porter and the library at their commits; wikiload under the old module path passes; another loader fails
package main

import (
	"testing"
	"time"

	"github.com/geul-org/wikiport/internal/port"
	"github.com/geul-org/wikiport/internal/source"
)

func TestProvenanceOf(t *testing.T) {
	d := source.Dump{URL: "u", SHA256: "s", DownloadedAt: time.Unix(5, 0), Loader: wikiloadPath, LoaderCommit: "l"}
	p, err := provenanceOf(d, "abc", "def")
	if err != nil || p.DumpURI != "u" || p.DumpSHA256 != "s" || p.AcquiredAt != 5 || len(p.Tools) != 3 ||
		p.Tools[0] != (port.Tool{Doc: 1, Commit: "l"}) || p.Tools[1] != (port.Tool{Doc: 1, Commit: "abc"}) || p.Tools[2] != (port.Tool{Doc: 2, Commit: "def"}) {
		t.Fatalf("got %+v %v", p, err)
	}
	d.Loader = legacyWikiloadPath
	if _, err := provenanceOf(d, "abc", "def"); err != nil {
		t.Fatalf("wikiload under the old module path: %v", err)
	}
	d.Loader = "github.com/geul-org/geul-entity/artifacts/wikidata"
	if _, err := provenanceOf(d, "abc", "def"); err == nil {
		t.Fatal("other loader: want error")
	}
}
