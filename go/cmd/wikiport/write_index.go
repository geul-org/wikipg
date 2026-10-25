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
// prefix+number for each written document and every document already
// published under the Issuer (publishedNumbers), so a run over another range
// does not drop the names of earlier runs (publish.NameIndex).
func writeIndex(root string, issuer []core.Word, prefix string, nums []uint64) error {
	dir := filepath.Join(root, ".well-known", "geul", core.FormatID(issuer))
	published, err := publishedNumbers(dir)
	if err != nil {
		return err
	}
	b, err := publish.NameIndex(issuer, prefix, append(published, nums...))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "documents.geul"), b, 0o644)
}
