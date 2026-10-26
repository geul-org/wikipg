//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what writeReport — the run report as TSV, one line per entity after a header
package main

import (
	"os"
	"strings"

	"github.com/geul-org/wikiport/go/internal/port"
)

// writeReport writes reps to path as TSV (reportLine) after a header line.
func writeReport(path string, reps []port.Report) error {
	lines := []string{"id\tlabels\taliases\ttriples\tqualifiers\tskipped\traw"}
	for _, r := range reps {
		lines = append(lines, reportLine(r))
	}
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}
