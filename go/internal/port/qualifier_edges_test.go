//ff:func feature=port type=encoder control=sequence
//ff:what TestQualifierEdges — convertible qualifiers become QUALIFIERs on the target in order; bad properties and values are counted with the qualifier prefix, raw-kept dates too
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestQualifierEdges(t *testing.T) {
	rep := Report{Skipped: map[string]int{}, Raw: map[string]int{}}
	next := uint64(10)
	rs := qualifierEdges(&next, 9, []source.Qualifier{
		{Property: "P580", Type: "time", Value: `{"time":"+1971-00-00T00:00:00Z","timezone":0,"before":0,"after":0,"precision":9,"calendarmodel":"` + gregorian + `"}`},
		{Property: "P512", Type: "wikibase-entityid", Value: `{"entity-type":"item","id":"Q1765120"}`},
		{Property: "X1", Type: "string", Value: "a"},
		{Property: "P1114", Type: "quantity", Value: "{}"},
		{Property: "P585", Type: "time", Value: "x"},
		{Property: "P642", Type: "wikibase-entityid", Value: `{"id":"P31"}`},
	}, &rep)
	if len(rs) != 4 || next != 14 || rep.Qualifiers != 4 || rs[0].Payload[1] != 10 || rs[3].Payload[1] != 13 {
		t.Fatalf("got %d next %d %+v", len(rs), next, rep)
	}
	if rep.Skipped["qualifier 속성 ID: X1"] != 1 || rep.Skipped["qualifier 형식 미지원: quantity"] != 1 || rep.Raw["qualifier time: JSON 아님"] != 1 {
		t.Fatalf("counts %v %v", rep.Skipped, rep.Raw)
	}
	if !slices.Equal(rep.Properties, []string{"P580", "P512", "P585", "P642", "P31"}) {
		t.Fatalf("properties %v", rep.Properties)
	}
}
