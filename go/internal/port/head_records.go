//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what headRecords — the HEAD of a ported document: SPEC-VERSION, DOC-ID, SOURCE (dump, entity), ACQUIRED-AT, one TRANSFORMED-BY per tool
package port

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// headRecords writes the HEAD of the document for entity subject (its
// Reference) numbered num under issuer: SPEC-VERSION, DOC-ID
// DOC-POINTER(issuer, num, 0), SOURCE twice (the dump URI, then the entity,
// each with its digest field left out as the Absent record),
// ACQUIRED-AT (when the dump file was received) and one TRANSFORMED-BY per
// tool in p.Tools: the tool's Reference (toolRef) and its commit
// (CommitDigest). A malformed commit is an error. No wall-clock time is
// written, so the same input and tools give the same document.
func headRecords(issuer []core.Word, num uint64, subject core.Record, p Provenance) ([]core.Record, error) {
	id, err := core.EncodeDocID(issuer, []core.Word{core.Word(num >> 48), core.Word(num >> 32), core.Word(num >> 16), core.Word(num)})
	if err != nil {
		return nil, err
	}
	absent := core.Record{Kind: core.KindReference}
	rs := []core.Record{
		core.EncodeSpecVersion(core.SpecVersion),
		id,
		compose(core.KindMeta, metaSource, core.EncodeBytesValue(codecURIEven, []byte(p.DumpURI)), absent),
		compose(core.KindMeta, metaSource, subject, absent),
		compose(core.KindMeta, metaAcquiredAt, core.EncodePosixSeconds(p.AcquiredAt)),
	}
	for _, t := range p.Tools {
		commit, err := CommitDigest(t.Commit)
		if err != nil {
			return nil, err
		}
		rs = append(rs, compose(core.KindMeta, metaTransformedBy, toolRef(t.Doc), commit))
	}
	return rs, nil
}
