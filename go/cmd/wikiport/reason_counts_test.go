//ff:func feature=cli type=helper control=sequence
//ff:what TestReasonCounts — sorted reason=count pairs; empty for no reasons
package main

import "testing"

func TestReasonCounts(t *testing.T) {
	if got := reasonCounts(map[string]int{"b": 1, "a": 2}); got != "a=2; b=1" {
		t.Fatalf("got %q", got)
	}
	if got := reasonCounts(nil); got != "" {
		t.Fatalf("empty: %q", got)
	}
}
