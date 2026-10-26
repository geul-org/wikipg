//ff:func feature=port type=encoder control=sequence
//ff:what TestTripleEdges — convertible claims become TRIPLEs, each followed by its QUALIFIERs; bad predicates, values and ranks are counted, raw-kept dates too; properties used are collected
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/wikiport/go/internal/source"
)

func TestTripleEdges(t *testing.T) {
	subj, _ := entityRef("Q1")
	rep := Report{Skipped: map[string]int{}, Raw: map[string]int{}}
	next := uint64(1)
	rs := tripleEdges(&next, subj, []source.Claim{
		{Property: "P31", Type: "wikibase-entityid", Value: "Q5", Rank: "normal", Qualifiers: []source.Qualifier{{Property: "P580", Type: "time", Value: "x"}}},
		{Property: "P1659", Type: "wikibase-entityid", Value: "P279", Rank: "preferred"},
		{Property: "X1", Type: "string", Value: "a", Rank: "normal"},
		{Property: "P2", Type: "unknown", Value: "{}", Rank: "normal"},
		{Property: "P3", Type: "string", Value: "a", Rank: "odd"},
		{Property: "P585", Type: "time", Value: "x", Rank: "normal"},
	}, &rep)
	if rep.Raw["time: JSON 아님"] != 1 || rep.Raw["qualifier time: JSON 아님"] != 1 || rep.Triples != 3 || rep.Qualifiers != 1 {
		t.Fatalf("raw %v", rep.Raw)
	}
	if len(rs) != 4 || next != 5 || rs[0].Codec != edgeTriple || rs[1].Codec != edgeQualifier || rs[1].Payload[5] != 1 || rs[2].Payload[len(rs[2].Payload)-1] != 1 {
		t.Fatalf("got %d next %d", len(rs), next)
	}
	if rep.Skipped["속성 ID: X1"] != 1 || rep.Skipped["형식 미지원: unknown"] != 1 || rep.Skipped["rank: odd"] != 1 {
		t.Fatalf("skipped %v", rep.Skipped)
	}
	if !slices.Equal(rep.Properties, []string{"P31", "P580", "P1659", "P279", "P585"}) {
		t.Fatalf("properties %v", rep.Properties)
	}
}
