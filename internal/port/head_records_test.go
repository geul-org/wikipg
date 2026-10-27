//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestHeadRecords — SPEC-VERSION, DOC-ID, two SOURCE (dump digest or left out, entity left out), ACQUIRED-AT and one TRANSFORMED-BY per tool in order; a bad Issuer or commit fails
package port

import (
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHeadRecords(t *testing.T) {
	subj, _ := entityRef("Q42")
	rs, err := headRecords([]core.Word{1, 0}, 42, subj, prov())
	if err != nil || len(rs) != 8 {
		t.Fatalf("got %d %v", len(rs), err)
	}
	var codecs []uint8
	for _, r := range rs {
		codecs = append(codecs, r.Codec)
	}
	if !slices.Equal(codecs, []uint8{core.MetaSpecVersion, core.MetaDocID, metaSource, metaSource, metaAcquiredAt, metaTransformedBy, metaTransformedBy, metaTransformedBy}) {
		t.Fatalf("codecs %02X", codecs)
	}
	if issuer, doc, err := core.DecodeDocID(rs[1]); err != nil || !slices.Equal(issuer, []core.Word{1, 0}) || doc[3] != 42 {
		t.Fatalf("DOC-ID: %v %v %v", issuer, doc, err)
	}
	porter := append([]core.Word{0x0360, 0, 1, 0, 0, 1, 2, 0x0D62}, slices.Repeat([]core.Word{0x2222}, 10)...)
	if !slices.Equal(rs[3].Payload, []core.Word{0x00A0, 42, core.Absent}) || rs[2].Payload[len(rs[2].Payload)-1] != core.Absent || !slices.Equal(rs[6].Payload, porter) {
		t.Fatalf("subject %04X porter %04X", rs[3].Payload, rs[6].Payload)
	}
	if _, err := headRecords([]core.Word{0, 1}, 42, subj, prov()); err == nil {
		t.Fatal("non-canonical Issuer: want error")
	}
	withSum := prov()
	withSum.DumpSHA256 = strings.Repeat("ab", 32)
	if rs, err := headRecords([]core.Word{1, 0}, 42, subj, withSum); err != nil || rs[2].Payload[len(rs[2].Payload)-16] != 0xabab {
		t.Fatalf("dump digest: %v", err)
	}
	withSum.DumpSHA256 = "ab"
	if _, err := headRecords([]core.Word{1, 0}, 42, subj, withSum); err == nil {
		t.Fatal("short dump digest: want error")
	}
	bad := prov()
	bad.Tools[1].Commit = strings.Repeat("2", 7)
	if _, err := headRecords([]core.Word{1, 0}, 42, subj, bad); err == nil {
		t.Fatal("short commit: want error")
	}
}
