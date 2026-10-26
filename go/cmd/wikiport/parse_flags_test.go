//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestParseFlags — defaults and overrides; a missing or short commit, a bad range and an unknown flag fail
package main

import (
	"slices"
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	a, b := strings.Repeat("a", 40), strings.Repeat("b", 40)
	c, err := parseFlags([]string{"-commit", a, "-lib-commit", b})
	if err != nil || c.Root != "." || c.From != 1 || c.To != 1000 || !slices.Equal(c.Langs, []string{"ko", "en"}) || c.Commit != a || c.LibCommit != b {
		t.Fatalf("defaults: %+v %v", c, err)
	}
	c, err = parseFlags([]string{"-commit", a, "-lib-commit", b, "-from", "5", "-to", "9", "-langs", "en, ko", "-report", "r.tsv", "-root", "/x"})
	if err != nil || c.From != 5 || c.To != 9 || !slices.Equal(c.Langs, []string{"en", "ko"}) || c.Report != "r.tsv" || c.Root != "/x" {
		t.Fatalf("overrides: %+v %v", c, err)
	}
	for _, args := range [][]string{{}, {"-commit", a}, {"-commit", "abc", "-lib-commit", b}, {"-commit", a, "-lib-commit", b, "-from", "0"}, {"-commit", a, "-lib-commit", b, "-from", "9", "-to", "5"}, {"-nope"}} {
		if _, err := parseFlags(args); err == nil {
			t.Fatalf("%v: want error", args)
		}
	}
}
