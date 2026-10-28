//ff:func feature=cli type=command control=sequence
//ff:what TestRunCommand — the repair-reference-snak subcommand and a plain load both dispatch (bad flags fail in each)
package main

import (
	"strings"
	"testing"
)

func TestRunCommand(t *testing.T) {
	if err := runCommand([]string{"repair-reference-snak", "-batch", "0"}); err == nil || !strings.Contains(err.Error(), "-batch") {
		t.Fatalf("repair: %v", err)
	}
	if err := runCommand([]string{"-in", "x"}); err == nil || !strings.Contains(err.Error(), "-commit") {
		t.Fatalf("load: %v", err)
	}
	if err := runCommand(nil); err == nil {
		t.Fatal("no args: want error")
	}
}
