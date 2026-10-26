//ff:func feature=cli type=helper control=sequence
//ff:what TestReportLine — tab-separated counts (labels, descriptions, aliases …), sorted skip reasons and raw-kept date reasons
package main

import (
	"testing"

	"github.com/geul-org/wikiport/go/internal/port"
)

func TestReportLine(t *testing.T) {
	got := reportLine(port.Report{ID: "Q1", Labels: 2, Descriptions: 6, Aliases: 3, Triples: 4, Qualifiers: 5, Skipped: map[string]int{"b": 1, "a": 2}, Raw: map[string]int{"time: x": 1}})
	if got != "Q1\t2\t6\t3\t4\t5\ta=2; b=1\ttime: x=1" {
		t.Fatalf("got %q", got)
	}
}
