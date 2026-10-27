//ff:func feature=port type=encoder control=sequence
//ff:what TestQualifierEdge — Edge Id, Property, the target as a LOCAL Reference, Value (the edge/QUALIFIER example)
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQualifierEdge(t *testing.T) {
	prop, _ := entityRef("P580")
	r := qualifierEdge(freshIDs(10), prop, 9, core.EncodeBytesValue(core.CodecEDTFEven, []byte("1971")))
	ws, _ := core.EncodeRecord(r)
	if r.Codec != edgeQualifier || !slices.Equal(ws, []core.Word{0x148B, 0x0088, 0x000A, 0x00A8, 0x0244, 0x0088, 0x0009, 0x0924, 0x3139, 0x3731}) {
		t.Fatalf("got %04X", ws)
	}
}
