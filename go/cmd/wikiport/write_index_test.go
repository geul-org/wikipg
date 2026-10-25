//ff:func feature=cli type=helper control=sequence
//ff:what TestWriteIndex — documents.geul under the Issuer directory; a bad Issuer and an unwritable root fail
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestWriteIndex(t *testing.T) {
	root := t.TempDir()
	if err := writeIndex(root, []core.Word{1, 1}, "p", []uint64{31}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".well-known/geul/0x00010001/documents.geul")); err != nil {
		t.Fatal(err)
	}
	if err := writeIndex(root, []core.Word{0, 1}, "p", []uint64{31}); err == nil {
		t.Fatal("bad Issuer: want error")
	}
	blocked := filepath.Join(root, "blocked")
	os.WriteFile(blocked, nil, 0o644)
	if err := writeIndex(blocked, []core.Word{1, 1}, "p", []uint64{31}); err == nil {
		t.Fatal("unwritable: want error")
	}
}
