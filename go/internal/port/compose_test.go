//ff:func feature=port type=encoder control=sequence
//ff:what TestCompose — the payload is each part's full encoding in order
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCompose(t *testing.T) {
	r := compose(core.KindMeta, metaSource, core.EncodeLocalRef(7), core.EncodeUintValue(5))
	if r.Kind != core.KindMeta || r.Codec != metaSource || !slices.Equal(r.Payload, []core.Word{0x0088, 7, 0x0882, 5}) {
		t.Fatalf("got %+v", r)
	}
	if r := compose(core.KindEdge, edgeLabel); len(r.Payload) != 0 {
		t.Fatalf("empty: %+v", r)
	}
}
