//ff:func feature=cli type=command control=sequence
//ff:what TestPortOne — Q42 is published under its number; a missing ID is false; an unwritable tree fails
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"

	"github.com/geul-org/wikiport/go/internal/source"
)

func TestPortOne(t *testing.T) {
	db, schema := fixtureDB(t)
	root := t.TempDir()
	c := config{Langs: []string{"ko", "en"}, Commit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LibCommit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Schema: schema, Dump: 1}
	prov, _ := provenanceOf(source.Dump{URL: "u", Loader: wikiloadPath, LoaderCommit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, c.Commit, c.LibCommit)
	rep, ok, err := portOne(db, resolve.Dir{Root: root}, "Q42", []core.Word{1, 0}, c, prov)
	if err != nil || !ok || rep.Triples == 0 || rep.Qualifiers == 0 {
		t.Fatalf("Q42: %+v %v %v", rep, ok, err)
	}
	if _, err := os.Stat(filepath.Join(root, ".well-known/geul/0x00010000/0x000000000000002a/0x00000001.geul")); err != nil {
		t.Fatal(err)
	}
	if _, ok, err := portOne(db, resolve.Dir{Root: root}, "Q6", []core.Word{1, 0}, c, prov); ok || err != nil {
		t.Fatalf("Q6: %v %v", ok, err)
	}
	blocked := filepath.Join(root, "blocked")
	os.WriteFile(blocked, nil, 0o644)
	if _, _, err := portOne(db, resolve.Dir{Root: blocked}, "Q42", []core.Word{1, 0}, c, prov); err == nil {
		t.Fatal("unwritable: want error")
	}
}
