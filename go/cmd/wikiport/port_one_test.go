//ff:func feature=cli type=command control=sequence
//ff:what TestPortOne — Q42 is published under its number; a missing ID is false; an unwritable tree fails
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
)

func TestPortOne(t *testing.T) {
	db := openDB(t)
	root := t.TempDir()
	c := config{Langs: []string{"ko", "en"}, Commit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LibCommit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	rep, ok, err := portOne(db, resolve.Dir{Root: root}, "Q42", []core.Word{1, 0}, c)
	if err != nil || !ok || rep.Triples == 0 || rep.Qualifiers == 0 {
		t.Fatalf("Q42: %+v %v %v", rep, ok, err)
	}
	if _, err := os.Stat(filepath.Join(root, ".well-known/geul/0x00010000/0x000000000000002a/0x00000001.geul")); err != nil {
		t.Fatal(err)
	}
	if _, ok, err := portOne(db, resolve.Dir{Root: root}, "Q0", []core.Word{1, 0}, c); ok || err != nil {
		t.Fatalf("Q0: %v %v", ok, err)
	}
	blocked := filepath.Join(root, "blocked")
	os.WriteFile(blocked, nil, 0o644)
	if _, _, err := portOne(db, resolve.Dir{Root: blocked}, "Q42", []core.Word{1, 0}, c); err == nil {
		t.Fatal("unwritable: want error")
	}
}
