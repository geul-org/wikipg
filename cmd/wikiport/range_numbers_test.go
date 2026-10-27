//ff:func feature=cli type=helper control=sequence
//ff:what TestRangeNumbers — inclusive range
package main

import (
	"slices"
	"testing"
)

func TestRangeNumbers(t *testing.T) {
	if got := rangeNumbers(3, 5); !slices.Equal(got, []uint64{3, 4, 5}) {
		t.Fatalf("got %v", got)
	}
}
