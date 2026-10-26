//ff:func feature=cli type=helper control=sequence
//ff:what TestWriteReport — a header and one line per report; the header has as many columns as a row
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/wikiport/go/internal/port"
)

func TestWriteReport(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r.tsv")
	if err := writeReport(p, []port.Report{{ID: "Q1"}, {ID: "Q2"}}); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(p)
	if lines := strings.Split(strings.TrimSpace(string(b)), "\n"); len(lines) != 3 || !strings.HasPrefix(lines[0], "id\t") || !strings.HasPrefix(lines[2], "Q2\t") ||
		strings.Count(lines[0], "\t") != strings.Count(lines[1], "\t") {
		t.Fatalf("got %q", b)
	}
}
