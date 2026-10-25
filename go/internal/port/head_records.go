//ff:func feature=port type=encoder control=sequence
//ff:what headRecords — the HEAD of a ported document: SPEC-VERSION, DOC-ID, SOURCE (dump, entity), ACQUIRED-AT, TRANSFORMED-BY (loader, porter)
package port

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// headRecords writes the HEAD of the document for entity subject (its
// Reference) numbered num under issuer: SPEC-VERSION, DOC-ID
// DOC-POINTER(issuer, num, 0), SOURCE twice (the dump URI, then the entity),
// ACQUIRED-AT (when the dump file was received) and TRANSFORMED-BY for the
// loader and this porter. TRANSFORMED-BY carries a URI Value and the version
// as UTF8-TEXT instead of a Reference (wiki/SPEC-PROPOSAL.md). No wall-clock
// time is written, so the same input and porter give the same document.
func headRecords(issuer []core.Word, num uint64, subject core.Record, p Provenance) ([]core.Record, error) {
	id, err := core.EncodeDocID(issuer, []core.Word{core.Word(num >> 48), core.Word(num >> 32), core.Word(num >> 16), core.Word(num)})
	if err != nil {
		return nil, err
	}
	uri := func(s string) core.Record { return core.EncodeBytesValue(codecURIEven, []byte(s)) }
	text := func(s string) core.Record { return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s)) }
	return []core.Record{
		core.EncodeSpecVersion(core.SpecVersion),
		id,
		compose(core.KindMeta, metaSource, uri(p.DumpURI)),
		compose(core.KindMeta, metaSource, subject),
		compose(core.KindMeta, metaAcquiredAt, core.EncodePosixSeconds(p.AcquiredAt)),
		compose(core.KindMeta, metaTransformedBy, uri(p.Loader.URI), text(p.Loader.Version)),
		compose(core.KindMeta, metaTransformedBy, uri(p.Porter.URI), text(p.Porter.Version)),
	}, nil
}
