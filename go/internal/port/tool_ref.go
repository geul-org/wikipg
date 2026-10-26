//ff:func feature=port type=encoder control=sequence
//ff:what toolRef — the EXTERNAL-ISSUER-1WORD Reference to LOCAL 2 of geul.org URL-defined Reference document n
package port

import "github.com/geul-org/geul/go/pkg/core"

// toolRef points at the tool that geul.org document 0x0001000000000000+n
// defines (GEUL-ORG URL 정의 Reference 문서: LOCAL 2 is the tool) as
// EXTERNAL-ISSUER-1WORD(0x0000, the document, 2).
func toolRef(n uint16) core.Record {
	return core.Record{Kind: core.KindReference, Codec: refExternalIssuer1Word << 3, Payload: []core.Word{0x0000, 0x0001, 0, 0, core.Word(n), 2}}
}
