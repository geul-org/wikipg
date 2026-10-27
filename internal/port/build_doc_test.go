//ff:func feature=port type=encoder control=sequence
//ff:what TestBuildDoc — a full document with LABEL, DESCRIPTION, ALIAS, a TRIPLE and its QUALIFIER passes schema with its DOC-ID (Q in the top byte) and counts; a property states its datatype; a bad ID and a bad Issuer fail
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/schema"

	"github.com/geul-org/wikiport/internal/source"
)

func TestBuildDoc(t *testing.T) {
	e := source.Entity{
		ID:           "Q42",
		Labels:       []source.Text{{Lang: "en", Value: "Douglas Adams"}},
		Descriptions: []source.Text{{Lang: "en", Value: "English writer"}},
		Aliases:      []source.Text{{Lang: "en", Value: "DNA"}},
		Claims: []source.Claim{
			{Property: "P31", Type: "wikibase-entityid", Value: "Q5", Rank: "normal", Qualifiers: []source.Qualifier{{Property: "P642", Type: "wikibase-entityid", Value: "Q1"}}},
			{Property: "P2", Type: "unknown", Rank: "normal", Value: "{}"},
		},
	}
	b, rep, err := BuildDoc(e, []core.Word{1, 0}, prov(), freshIDs(1))
	if err != nil || rep.Labels != 1 || rep.Descriptions != 1 || rep.Aliases != 1 || rep.Triples != 1 || rep.Qualifiers != 1 || rep.Skipped["형식 미지원: unknown"] != 1 {
		t.Fatalf("report %+v %v", rep, err)
	}
	if fs := schema.Check(b, schema.Options{}).Facts; len(fs) != 0 {
		t.Fatalf("schema: %+v", fs)
	}
	ws, _ := core.ReadWords(b)
	if parts, err := core.SplitDocument(ws); err != nil || len(parts.Head) < 2 {
		t.Fatalf("split: %v", err)
	} else if _, doc, err := core.DecodeDocID(parts.Head[1]); err != nil || doc[0] != 0x5100 || doc[3] != 42 {
		t.Fatalf("DOC-ID: %04X %v", doc, err)
	}
	d, _ := inspect.Decode(b)
	if len(d.Body) != 5 || d.Body[1].Record.Codec != edgeDescription || d.Body[2].Record.Codec != edgeAlias || d.Body[3].ID != 4 || d.Body[4].ID != 5 {
		t.Fatalf("body %+v", d.Body)
	}
	pb, prep, err := BuildDoc(source.Entity{ID: "P214", Datatype: "external-id"}, []core.Word{1, 0}, prov(), freshIDs(1))
	if fs := schema.Check(pb, schema.Options{}).Facts; err != nil || len(fs) != 0 || len(prep.Skipped) != 0 {
		t.Fatalf("property: %v %+v", err, fs)
	}
	if pd, _ := inspect.Decode(pb); len(pd.Body) != 2 || pd.Body[1].Record.Codec != edgeTriple {
		t.Fatalf("property body %+v", pd.Body)
	}
	if _, _, err := BuildDoc(source.Entity{ID: "Q72057594037927936"}, []core.Word{1, 0}, prov(), freshIDs(1)); err == nil {
		t.Fatal("Q2^56: want error")
	}
	if _, _, err := BuildDoc(source.Entity{ID: "L1"}, []core.Word{1, 0}, prov(), freshIDs(1)); err == nil {
		t.Fatal("lexeme: want error")
	}
	if _, _, err := BuildDoc(e, []core.Word{0, 1}, prov(), freshIDs(1)); err == nil {
		t.Fatal("bad Issuer: want error")
	}
}
