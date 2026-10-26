//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what digestWords — digest bytes as WORDs, two octets each with the first high (SHA-256-DIGEST framing)
package port

import "github.com/geul-org/geul/go/pkg/core"

// digestWords frames an even-length digest b as WORDs: octets 2i and 2i+1
// become WORD i, the first octet high (value/binary/SHA-256-DIGEST).
func digestWords(b []byte) []core.Word {
	ws := make([]core.Word, len(b)/2)
	for i := range ws {
		ws[i] = core.Word(b[2*i])<<8 | core.Word(b[2*i+1])
	}
	return ws
}
