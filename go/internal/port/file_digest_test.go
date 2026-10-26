//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestFileDigest — 64 hex digits become SHA-256-DIGEST; empty is Absent; short or non-hex fails
package port

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestFileDigest(t *testing.T) {
	r, err := fileDigest(strings.Repeat("a2", 32))
	if err != nil || r.Codec != core.CodecSHA256Digest || len(r.Payload) != 16 || r.Payload[0] != 0xa2a2 {
		t.Fatalf("sum: %+v %v", r, err)
	}
	if r, err := fileDigest(""); err != nil || !core.IsAbsent(r) {
		t.Fatalf("empty: %+v %v", r, err)
	}
	for _, s := range []string{"abc", strings.Repeat("zz", 32)} {
		if _, err := fileDigest(s); err == nil {
			t.Fatalf("%q: want error", s)
		}
	}
}
