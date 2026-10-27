//ff:func feature=port type=helper control=sequence
//ff:what DocNumber — the document ID of a Wikidata entity under wikidata.geul.org: the prefix letter in the top byte, the number below (Q42 → 0x51000000_0000002A)
package port

import "fmt"

// DocNumber is the 4-word document ID of the Wikidata entity prefix+n under
// the one Issuer wikidata.geul.org (README 문서 ID): the ASCII letter
// (Q 0x51, P 0x50, L 0x4C …) in the top byte and n in the 56 bits below, so
// the hex reads as the kind. n must be below 2^56.
func DocNumber(prefix byte, n uint64) (uint64, error) {
	if n >= 1<<56 {
		return 0, fmt.Errorf("wikiport: %c%d exceeds the 56-bit document number", prefix, n)
	}
	return uint64(prefix)<<56 | n, nil
}
