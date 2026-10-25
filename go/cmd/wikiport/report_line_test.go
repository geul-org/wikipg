//ff:func feature=cli type=helper control=sequence
//ff:what TestReportLine — tab-separated counts and sorted skip reasons
package main

import (
	"testing"

	"github.com/geul-org/geul-dev/wiki/go/internal/port"
)

func TestReportLine(t *testing.T) {
	got := reportLine(port.Report{ID: "Q1", Labels: 2, Aliases: 3, Statements: 4, Skipped: map[string]int{"b": 1, "a": 2}})
	if got != "Q1\t2\t3\t4\ta=2; b=1" {
		t.Fatalf("got %q", got)
	}
}
