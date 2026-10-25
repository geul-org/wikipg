//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what reportLine — one TSV line of a report: ID, label, alias and statement counts, skipped reasons with counts
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/geul-org/geul-dev/wiki/go/internal/port"
)

// reportLine formats r as ID, labels, aliases, statements and the skipped
// reasons ("reason=count" joined by "; ", reasons sorted), tab-separated.
func reportLine(r port.Report) string {
	var reasons []string
	for reason, n := range r.Skipped {
		reasons = append(reasons, fmt.Sprintf("%s=%d", reason, n))
	}
	slices.Sort(reasons)
	return fmt.Sprintf("%s\t%d\t%d\t%d\t%s", r.ID, r.Labels, r.Aliases, r.Statements, strings.Join(reasons, "; "))
}
