//ff:func feature=port type=encoder control=sequence
//ff:what TestDigestWords — octet pairs become big-endian WORDs; empty input gives none
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDigestWords(t *testing.T) {
	if got := digestWords([]byte{0xdd, 0xc0, 0x71, 0x99}); !slices.Equal(got, []core.Word{0xddc0, 0x7199}) {
		t.Fatalf("got %04X", got)
	}
	if got := digestWords(nil); len(got) != 0 {
		t.Fatalf("empty: %04X", got)
	}
}
