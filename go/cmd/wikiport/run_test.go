//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRun — a small range writes both Issuers' documents, indexes and the report; bad flags, an unwritable root and unwritable Issuer directories fail
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	_, schema := fixtureDB(t)
	root := t.TempDir()
	report := filepath.Join(root, "r.tsv")
	if err := run([]string{"-root", root, "-from", "1", "-to", "3", "-langs", "en", "-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-lib-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-schema", schema, "-report", report}); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{".well-known/geul/0x00010000/documents.geul", ".well-known/geul/0x00010001/documents.geul", "r.tsv"} {
		if _, err := os.Stat(filepath.Join(root, p)); err != nil {
			t.Fatal(err)
		}
	}
	if err := run([]string{"-root", root, "-from", "1", "-to", "1", "-langs", "en", "-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-lib-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-schema", schema}); err != nil {
		t.Fatalf("no report: %v", err)
	}
	if err := run(nil); err == nil {
		t.Fatal("no commit: want error")
	}
	for _, dump := range []string{"9", "2"} {
		if err := run([]string{"-root", root, "-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-lib-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-schema", schema, "-dump-id", dump}); err == nil {
			t.Fatalf("dump %s (missing, other loader): want error", dump)
		}
	}
	if err := run([]string{"-root", root, "-dsn", "host=/nonexistent", "-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-lib-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-schema", schema}); err == nil {
		t.Fatal("unreachable database: want error")
	}
	blocked := filepath.Join(root, "blocked")
	os.WriteFile(blocked, nil, 0o644)
	if err := run([]string{"-root", blocked, "-from", "1", "-to", "1", "-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-lib-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-schema", schema}); err == nil {
		t.Fatal("unwritable: want error")
	}
	// An Issuer directory that is a file: the properties of Q1 cannot be
	// published, and with no documents (Q6 was never created) neither index
	// can be written.
	for _, c := range []struct{ file, from string }{
		{"0x00010001", "1"},
		{"0x00010000", "6"},
		{"0x00010001", "6"},
	} {
		r := t.TempDir()
		os.MkdirAll(filepath.Join(r, ".well-known/geul"), 0o755)
		os.WriteFile(filepath.Join(r, ".well-known/geul", c.file), nil, 0o644)
		if err := run([]string{"-root", r, "-from", c.from, "-to", c.from, "-langs", "en", "-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-lib-commit", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-schema", schema}); err == nil {
			t.Fatalf("%s file, Q%s: want error", c.file, c.from)
		}
	}
}
