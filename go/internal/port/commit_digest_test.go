//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestCommitDigest — a SHA-1 ID is the docs' SHA1-DIGEST, a SHA-256 ID a SHA-256-DIGEST; short, odd and non-hex IDs fail
package port

import (
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCommitDigest(t *testing.T) {
	r, err := CommitDigest("ddc0719974a32aaa7b422ee84221a9721085ce4c")
	ws, _ := core.EncodeRecord(r)
	if err != nil || !slices.Equal(ws, []core.Word{0x0D62, 0xddc0, 0x7199, 0x74a3, 0x2aaa, 0x7b42, 0x2ee8, 0x4221, 0xa972, 0x1085, 0xce4c}) {
		t.Fatalf("sha1: %04X %v", ws, err)
	}
	if r, err := CommitDigest(strings.Repeat("ab", 32)); err != nil || r.Codec != core.CodecSHA256Digest || len(r.Payload) != 16 || r.Payload[0] != 0xabab {
		t.Fatalf("sha256: %+v %v", r, err)
	}
	for _, id := range []string{"ddc0719", "abc", "zz" + strings.Repeat("0", 38), ""} {
		if _, err := CommitDigest(id); err == nil {
			t.Fatalf("%q: want error", id)
		}
	}
}
