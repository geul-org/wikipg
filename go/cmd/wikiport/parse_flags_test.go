//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestParseFlags — defaults and overrides; a missing commit, a bad range and an unknown flag fail
package main

import (
	"slices"
	"testing"
)

func TestParseFlags(t *testing.T) {
	c, err := parseFlags([]string{"-commit", "abc"})
	if err != nil || c.Root != "." || c.From != 1 || c.To != 1000 || !slices.Equal(c.Langs, []string{"ko", "en"}) || c.Commit != "abc" {
		t.Fatalf("defaults: %+v %v", c, err)
	}
	c, err = parseFlags([]string{"-commit", "abc", "-from", "5", "-to", "9", "-langs", "en, ko", "-report", "r.tsv", "-root", "/x"})
	if err != nil || c.From != 5 || c.To != 9 || !slices.Equal(c.Langs, []string{"en", "ko"}) || c.Report != "r.tsv" || c.Root != "/x" {
		t.Fatalf("overrides: %+v %v", c, err)
	}
	for _, args := range [][]string{{}, {"-commit", "a", "-from", "0"}, {"-commit", "a", "-from", "9", "-to", "5"}, {"-nope"}} {
		if _, err := parseFlags(args); err == nil {
			t.Fatalf("%v: want error", args)
		}
	}
}
