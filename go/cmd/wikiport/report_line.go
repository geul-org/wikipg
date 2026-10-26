//ff:func feature=cli type=helper control=sequence
//ff:what reportLine — one TSV line of a report: ID, label, alias, triple and qualifier counts, skipped and raw-kept reasons with counts
package main

import (
	"fmt"

	"github.com/geul-org/geul-dev/wiki/go/internal/port"
)

// reportLine formats r as ID, labels, aliases, triples, qualifiers and the skipped
// reasons, then the reasons of dates kept as original text (reasonCounts),
// tab-separated.
func reportLine(r port.Report) string {
	return fmt.Sprintf("%s\t%d\t%d\t%d\t%d\t%s\t%s", r.ID, r.Labels, r.Aliases, r.Triples, r.Qualifiers, reasonCounts(r.Skipped), reasonCounts(r.Raw))
}
