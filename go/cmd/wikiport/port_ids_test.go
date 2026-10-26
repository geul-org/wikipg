//ff:func feature=cli type=command control=sequence
//ff:what TestPortIDs — present IDs are written and missing ones only reported; the first error stops the run
package main

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
)

func TestPortIDs(t *testing.T) {
	db := openDB(t)
	c := config{Langs: []string{"en"}, Commit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LibCommit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	reps, written, err := portIDs(db, resolve.Dir{Root: t.TempDir()}, "Q", []uint64{1, 6}, []core.Word{1, 0}, c)
	if err != nil || len(reps) != 2 || !slices.Equal(written, []uint64{1}) {
		t.Fatalf("got %d %v %v", len(reps), written, err)
	}
	blocked := filepath.Join(t.TempDir(), "blocked")
	os.WriteFile(blocked, nil, 0o644)
	if _, _, err := portIDs(db, resolve.Dir{Root: blocked}, "Q", []uint64{1}, []core.Word{1, 0}, c); err == nil {
		t.Fatal("unwritable: want error")
	}
}
