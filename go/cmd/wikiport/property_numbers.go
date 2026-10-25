//ff:func feature=cli type=helper control=iteration dimension=2
//ff:what propertyNumbers — the distinct property numbers the reports used, ascending
package main

import (
	"slices"
	"strconv"

	"github.com/geul-org/geul-dev/wiki/go/internal/port"
)

// propertyNumbers collects the numbers of the properties (P…) the reports
// used as predicates or values, without duplicates, ascending.
func propertyNumbers(reps []port.Report) []uint64 {
	var out []uint64
	for _, r := range reps {
		for _, p := range r.Properties {
			if n, err := strconv.ParseUint(p[1:], 10, 64); err == nil {
				out = append(out, n)
			}
		}
	}
	slices.Sort(out)
	return slices.Compact(out)
}
