//ff:func feature=cli type=helper control=sequence
//ff:what provenanceOf — the provenance every document of this run records (wiki/README 원본 판본)
package main

import "github.com/geul-org/wikiport/go/internal/port"

// provenanceOf is the provenance of the geuldev Wikidata load (wiki/README
// 원본 판본): the 20250908 JSON dump, received 2025-09-17T11:49:51Z, and the
// tools as geul.org URL-defined Reference documents (GEUL-ORG) — loaded by
// geul-entity (document 3) at 4225b3b, ported by wikiport (document 1) at
// commit with the GEUL library (document 2) at libCommit.
func provenanceOf(commit, libCommit string) port.Provenance {
	return port.Provenance{
		DumpURI:    "https://dumps.wikimedia.org/wikidatawiki/entities/20250908/wikidata-20250908-all.json.bz2",
		AcquiredAt: 1758109791,
		Tools: []port.Tool{
			{Doc: 3, Commit: "4225b3b4297796a79e95c4cda5aa54df3acc4069"},
			{Doc: 1, Commit: commit},
			{Doc: 2, Commit: libCommit},
		},
	}
}
