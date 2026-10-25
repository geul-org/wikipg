//ff:func feature=port type=encoder control=sequence
//ff:what TestStatementEdges — convertible claims become STATEMENTs; bad predicates, values and ranks are counted, raw-kept dates too; properties used are collected
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestStatementEdges(t *testing.T) {
	subj, _ := entityRef("Q1")
	rep := Report{Skipped: map[string]int{}, Raw: map[string]int{}}
	next := uint64(1)
	rs := statementEdges(&next, subj, []source.Claim{
		{Property: "P31", Type: "wikibase-entityid", Value: "Q5", Rank: "normal"},
		{Property: "P1659", Type: "wikibase-entityid", Value: "P279", Rank: "preferred"},
		{Property: "X1", Type: "string", Value: "a", Rank: "normal"},
		{Property: "P2", Type: "quantity", Value: "{}", Rank: "normal"},
		{Property: "P3", Type: "string", Value: "a", Rank: "odd"},
		{Property: "P585", Type: "time", Value: "x", Rank: "normal"},
	}, &rep)
	if rep.Raw["time: JSON 아님"] != 1 {
		t.Fatalf("raw %v", rep.Raw)
	}
	if len(rs) != 3 || next != 4 || rs[1].Payload[len(rs[1].Payload)-1] != 1 {
		t.Fatalf("got %d next %d", len(rs), next)
	}
	if rep.Skipped["속성 ID: X1"] != 1 || rep.Skipped["형식 미지원: quantity"] != 1 || rep.Skipped["rank: odd"] != 1 {
		t.Fatalf("skipped %v", rep.Skipped)
	}
	if !slices.Equal(rep.Properties, []string{"P31", "P1659", "P279", "P585"}) {
		t.Fatalf("properties %v", rep.Properties)
	}
}
