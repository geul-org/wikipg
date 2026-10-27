//ff:func feature=port type=encoder control=sequence
//ff:what TestToolRef — wikiport (document 1) encodes as the GEUL-ORG table's Reference
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestToolRef(t *testing.T) {
	ws, err := core.EncodeRecord(toolRef(1))
	if err != nil || !slices.Equal(ws, []core.Word{0x0360, 0x0000, 0x0001, 0, 0, 0x0001, 0x0002}) {
		t.Fatalf("got %04X %v", ws, err)
	}
}
