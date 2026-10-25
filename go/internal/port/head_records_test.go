//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestHeadRecords — SPEC-VERSION, DOC-ID, two SOURCE, ACQUIRED-AT and two TRANSFORMED-BY in order; a bad Issuer fails
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHeadRecords(t *testing.T) {
	subj, _ := entityRef("Q42")
	rs, err := headRecords([]core.Word{1, 0}, 42, subj, prov())
	if err != nil || len(rs) != 7 {
		t.Fatalf("got %d %v", len(rs), err)
	}
	var codecs []uint8
	for _, r := range rs {
		codecs = append(codecs, r.Codec)
	}
	if !slices.Equal(codecs, []uint8{core.MetaSpecVersion, core.MetaDocID, metaSource, metaSource, metaAcquiredAt, metaTransformedBy, metaTransformedBy}) {
		t.Fatalf("codecs %02X", codecs)
	}
	if issuer, doc, err := core.DecodeDocID(rs[1]); err != nil || !slices.Equal(issuer, []core.Word{1, 0}) || doc[3] != 42 {
		t.Fatalf("DOC-ID: %v %v %v", issuer, doc, err)
	}
	if !slices.Equal(rs[3].Payload, []core.Word{0x00A0, 42}) || !slices.Equal(rs[6].Payload, []core.Word{0x08C5, 0x7000, 0x08C1, 0x3200}) {
		t.Fatalf("subject %04X porter %04X", rs[3].Payload, rs[6].Payload)
	}
	if _, err := headRecords([]core.Word{0, 1}, 42, subj, prov()); err == nil {
		t.Fatal("non-canonical Issuer: want error")
	}
}
