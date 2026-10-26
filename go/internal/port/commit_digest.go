//ff:func feature=port type=encoder control=selection
//ff:what CommitDigest — a full git object ID in hex as SHA1-DIGEST (40 digits) or SHA-256-DIGEST (64 digits); an error otherwise
package port

import (
	"encoding/hex"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// CommitDigest encodes the git object ID id, full and in hex, as the Version
// of TRANSFORMED-BY: SHA1-DIGEST for a SHA-1 repository (40 digits),
// SHA-256-DIGEST for a SHA-256 one (64 digits). An abbreviated or malformed
// ID is an error.
func CommitDigest(id string) (core.Record, error) {
	b, err := hex.DecodeString(id)
	codec := core.CodecSHA1Digest
	switch {
	case err != nil:
		return core.Record{}, fmt.Errorf("wikiport: commit %q is not hex", id)
	case len(b) == 32:
		codec = core.CodecSHA256Digest
	case len(b) != 20:
		return core.Record{}, fmt.Errorf("wikiport: commit %q is not a full SHA-1 or SHA-256 object ID", id)
	}
	return core.Record{Kind: core.KindValue, Codec: codec, Payload: digestWords(b)}, nil
}
