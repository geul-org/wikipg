//ff:func feature=port type=encoder control=sequence
//ff:what TestBuildDoc — a full document passes schema with its DOC-ID and counts; a bad ID and a bad Issuer fail
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
	"github.com/geul-org/geul/go/pkg/schema"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestBuildDoc(t *testing.T) {
	e := source.Entity{
		ID:      "Q42",
		Labels:  []source.Text{{Lang: "en", Value: "Douglas Adams"}},
		Aliases: []source.Text{{Lang: "en", Value: "DNA"}},
		Claims:  []source.Claim{{Property: "P31", Type: "wikibase-entityid", Value: "Q5", Rank: "normal"}, {Property: "P2", Type: "quantity", Rank: "normal", Value: "{}"}},
	}
	b, rep, err := BuildDoc(e, []core.Word{1, 0}, prov())
	if err != nil || rep.Labels != 1 || rep.Aliases != 1 || rep.Statements != 1 || rep.Skipped["형식 미지원: quantity"] != 1 {
		t.Fatalf("report %+v %v", rep, err)
	}
	if fs := schema.Check(b, schema.Options{}).Facts; len(fs) != 0 {
		t.Fatalf("schema: %+v", fs)
	}
	d, _ := inspect.Decode(b)
	if len(d.Body) != 3 || d.Body[2].ID != 3 {
		t.Fatalf("body %+v", d.Body)
	}
	if _, _, err := BuildDoc(source.Entity{ID: "L1"}, []core.Word{1, 0}, prov()); err == nil {
		t.Fatal("lexeme: want error")
	}
	if _, _, err := BuildDoc(e, []core.Word{0, 1}, prov()); err == nil {
		t.Fatal("bad Issuer: want error")
	}
}
