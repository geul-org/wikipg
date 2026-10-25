//ff:func feature=cli type=helper control=sequence
//ff:what writeIndex — the Issuer's documents.geul for the documents written
package main

import (
	"os"
	"path/filepath"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/publish"
)

// writeIndex writes `<root>/.well-known/geul/<issuer>/documents.geul` naming
// each written document prefix+number (publish.NameIndex).
func writeIndex(root string, issuer []core.Word, prefix string, nums []uint64) error {
	b, err := publish.NameIndex(issuer, prefix, nums)
	if err != nil {
		return err
	}
	dir := filepath.Join(root, ".well-known", "geul", core.FormatID(issuer))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "documents.geul"), b, 0o644)
}
