//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what edgeWith — an Edge whose Edge Id comes from the carried-over assignment for its content: records, then raw tail words
package port

import "github.com/geul-org/geul/go/pkg/core"

// edgeWith builds an Edge of codec with parts (whole records) and tail (raw
// words, such as TRIPLE's rank) after its Edge Id, which a asks for by that
// content (IDs.Edge), and returns it with the ID.
func edgeWith(a IDs, codec uint8, parts []core.Record, tail ...core.Word) (core.Record, uint64) {
	var rest []core.Word
	for _, p := range parts {
		ws, _ := core.EncodeRecord(p)
		rest = append(rest, ws...)
	}
	rest = append(rest, tail...)
	id := a.Edge(codec, rest)
	lead, _ := core.EncodeRecord(core.EncodeLocalRef(id))
	return core.Record{Kind: core.KindEdge, Codec: codec, Payload: append(lead, rest...)}, id
}
