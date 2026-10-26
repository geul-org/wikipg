//ff:func feature=port type=encoder control=selection
//ff:what fileDigest — a file's SHA-256 in hex as the SHA-256-DIGEST of SOURCE; the Absent record when unknown; an error when malformed
package port

import (
	"encoding/hex"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// fileDigest encodes sum, a SHA-256 in hex (64 digits), as the digest field
// of SOURCE (meta/provenance/SOURCE): a SHA-256-DIGEST Value. An empty sum is
// the Absent record (digest left out); any other form is an error.
func fileDigest(sum string) (core.Record, error) {
	b, err := hex.DecodeString(sum)
	switch {
	case sum == "":
		return core.Record{Kind: core.KindReference}, nil
	case err != nil || len(b) != 32:
		return core.Record{}, fmt.Errorf("wikiport: dump SHA-256 %q is not 64 hex digits", sum)
	}
	return core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: digestWords(b)}, nil
}
