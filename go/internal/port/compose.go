//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what compose — a record whose payload is the given records one after another
package port

import "github.com/geul-org/geul/go/pkg/core"

// compose builds a record of kind k and codec whose payload is parts, each
// encoded whole (TAG and Length included), in order.
func compose(k core.Kind, codec uint8, parts ...core.Record) core.Record {
	var payload []core.Word
	for _, p := range parts {
		ws, _ := core.EncodeRecord(p)
		payload = append(payload, ws...)
	}
	return core.Record{Kind: k, Codec: codec, Payload: payload}
}
