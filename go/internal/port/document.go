//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what document — a GEUL document of HEAD records, BODY, the BODY records and END, wrapped in DOC
package port

import "github.com/geul-org/geul/go/pkg/core"

// document lays out head, the BODY marker, body and the END marker, and wraps
// them in DOC (core.WrapDocument).
func document(head, body []core.Record) []byte {
	var in []core.Word
	for _, r := range head {
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	in = append(in, core.Word(0x180E))
	for _, r := range body {
		ws, _ := core.EncodeRecord(r)
		in = append(in, ws...)
	}
	return core.PutWords(core.WrapDocument(append(in, core.Word(0x180F))))
}
