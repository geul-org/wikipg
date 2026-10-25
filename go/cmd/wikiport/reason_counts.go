//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what reasonCounts — "reason=count" pairs joined by "; ", reasons sorted
package main

import (
	"fmt"
	"slices"
	"strings"
)

// reasonCounts formats counts as "reason=count" joined by "; ", reasons
// sorted so the report is the same on every run.
func reasonCounts(counts map[string]int) string {
	var out []string
	for reason, n := range counts {
		out = append(out, fmt.Sprintf("%s=%d", reason, n))
	}
	slices.Sort(out)
	return strings.Join(out, "; ")
}
