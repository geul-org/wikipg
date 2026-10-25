//ff:func feature=port type=encoder control=selection
//ff:what entityRef — the GEUL Reference of a Wikidata item or property ID (WIKIDATA-QID, WIKIDATA-PID); false for other IDs
package port

import "github.com/geul-org/geul/go/pkg/core"

// entityRef returns the direct Reference of a Wikidata item (Q → WIKIDATA-QID)
// or property (P → WIKIDATA-PID), Quantifier Identified, the number in the
// fewest words. Lexemes, forms, senses and entity schemas have no Reference
// codec and are false.
func entityRef(id string) (core.Record, bool) {
	prefix, n, ok := entityNumber(id)
	if !ok {
		return core.Record{}, false
	}
	words := core.EncodeUintValue(n).Payload
	switch prefix {
	case 'Q':
		return core.Record{Kind: core.KindReference, Codec: core.RefWikidataQID << 3, Payload: words}, true
	case 'P':
		return core.Record{Kind: core.KindReference, Codec: core.RefWikidataPID << 3, Payload: words}, true
	default:
		return core.Record{}, false
	}
}
