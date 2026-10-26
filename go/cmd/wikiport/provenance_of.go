//ff:func feature=cli type=helper control=sequence
//ff:what provenanceOf — the provenance every document of this run records (wiki/README 원본 판본)
package main

import "github.com/geul-org/wikiport/go/internal/port"

// provenanceOf is the provenance of the geuldev Wikidata load (wiki/README
// 원본 판본): the 20250908 JSON dump, received 2025-09-17T11:49:51Z, loaded by
// geul-entity's postgres.go at 4225b3b, ported by this tool at commit.
func provenanceOf(commit string) port.Provenance {
	return port.Provenance{
		DumpURI:    "https://dumps.wikimedia.org/wikidatawiki/entities/20250908/wikidata-20250908-all.json.bz2",
		AcquiredAt: 1758109791,
		Loader:     port.Tool{URI: "https://github.com/geul-org/geul-entity/blob/4225b3b4297796a79e95c4cda5aa54df3acc4069/artifacts/wikidata/postgres.go", Version: "4225b3b4297796a79e95c4cda5aa54df3acc4069"},
		Porter:     port.Tool{URI: "https://github.com/geul-org/wikiport/tree/main/go", Version: commit},
	}
}
