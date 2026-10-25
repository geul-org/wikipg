//ff:func feature=cli type=helper control=sequence
//ff:what TestWriteIndex — documents.geul under the Issuer directory names this run's and earlier published documents; a bad Issuer and an unwritable root fail
package main

import (
	"bytes"
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
	os.Mkdir(filepath.Join(root, ".well-known/geul/0x00010001/0x0000000000000007"), 0o755)
	if err := writeIndex(root, []core.Word{1, 1}, "p", []uint64{31}); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(root, ".well-known/geul/0x00010001/documents.geul"))
	if err != nil || !bytes.Contains(b, []byte("p7")) || !bytes.Contains(b, []byte("p31")) {
		t.Fatalf("names: %v", err)
	}
	if err := writeIndex(root, []core.Word{0, 1}, "p", []uint64{31}); err == nil {
		t.Fatal("bad Issuer: want error")
	}
	blocked := filepath.Join(root, "blocked")
	os.WriteFile(blocked, nil, 0o644)
	os.MkdirAll(filepath.Join(root, "unreadable/.well-known/geul"), 0o755)
	os.WriteFile(filepath.Join(root, "unreadable/.well-known/geul/0x00010001"), nil, 0o644)
	if err := writeIndex(filepath.Join(root, "unreadable"), []core.Word{1, 1}, "p", nil); err == nil {
		t.Fatal("Issuer path is a file: want error")
	}
	ro := filepath.Join(root, "ro/.well-known/geul")
	os.MkdirAll(ro, 0o755)
	os.Chmod(ro, 0o555)
	defer os.Chmod(ro, 0o755)
	if err := writeIndex(filepath.Join(root, "ro"), []core.Word{1, 1}, "p", nil); err == nil {
		t.Fatal("read-only parent: want error")
	}
	if err := writeIndex(blocked, []core.Word{1, 1}, "p", []uint64{31}); err == nil {
		t.Fatal("unwritable: want error")
	}
}
