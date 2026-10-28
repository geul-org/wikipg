//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestWriteReport — the report names the range and mode, the verdicts, the statement_reference check, errors and examples
package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/wikipg/internal/repair"
)

func TestWriteReport(t *testing.T) {
	var b bytes.Buffer
	st := repair.Stats{Spans: 1, Planned: 1, Hashes: repair.Tally{Content: 2}, Occurrences: repair.Tally{Order: 3}, RefDiffs: 1, RefDiffEntities: []string{"Q3"}, Errors: []string{"line 5"}, Examples: []string{"hx Q1 content"}}
	writeReport(&b, repair.Config{Kind: "Q", Lo: 1, Hi: 9, DryRun: true, DumpID: 1}, st)
	s := b.String()
	for _, want := range []string{"entities:Q1-Q9 (dry run)", "content 2", "order 3", "1 differ [Q3]", "line 5", "hx Q1 content"} {
		if !strings.Contains(s, want) {
			t.Fatalf("missing %q in %s", want, s)
		}
	}
	b.Reset()
	writeReport(&b, repair.Config{}, repair.Stats{})
	if !strings.Contains(b.String(), "(write)") {
		t.Fatal(b.String())
	}
}
