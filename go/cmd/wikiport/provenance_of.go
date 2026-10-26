//ff:func feature=cli type=helper control=sequence
//ff:what provenanceOf — the provenance every document of this run records, from the loaded dump's row (URL, SHA-256, time, loader) and the porter and library commits
package main

import (
	"fmt"

	"github.com/geul-org/wikiport/go/internal/port"
	"github.com/geul-org/wikiport/go/internal/source"
)

// wikiloadPath is the loader that wikiload records in the dump row; it lives
// in this repository, geul.org URL-defined Reference document 1 (wikiport).
const wikiloadPath = "github.com/geul-org/wikiport/go/cmd/wikiload"

// provenanceOf is the provenance of dump d (wiki/README 출처 기록): its URL,
// file SHA-256 and download time, then the tools as geul.org URL-defined Reference
// documents (GEUL-ORG) — loaded by wikiload (document 1) at d's loader commit,
// ported by wikiport (document 1) at commit with the GEUL library (document
// 2) at libCommit. A loader other than wikiload is an error: its Reference
// document is not known.
func provenanceOf(d source.Dump, commit, libCommit string) (port.Provenance, error) {
	if d.Loader != wikiloadPath {
		return port.Provenance{}, fmt.Errorf("wikiport: dump loaded by %q, not %s", d.Loader, wikiloadPath)
	}
	return port.Provenance{
		DumpURI:    d.URL,
		DumpSHA256: d.SHA256,
		AcquiredAt: d.DownloadedAt.Unix(),
		Tools: []port.Tool{
			{Doc: 1, Commit: d.LoaderCommit},
			{Doc: 1, Commit: commit},
			{Doc: 2, Commit: libCommit},
		},
	}, nil
}
