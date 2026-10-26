//ff:func feature=port type=encoder control=sequence
//ff:what TestDatatypeEdges — a property gets REFERENCE-NODE(propertyType URI) and TRIPLE(P, ExternalId URI) with three ids used; an item gets none; an unknown datatype is skipped
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDatatypeEdges(t *testing.T) {
	subject, _ := entityRef("P214")
	rep := Report{Skipped: map[string]int{}}
	next := uint64(4)
	rs := datatypeEdges(&next, subject, "external-id", &rep)
	if len(rs) != 2 || next != 7 || rs[0].Codec != edgeReferenceNode || rs[1].Codec != edgeTriple {
		t.Fatalf("got %+v next %d", rs, next)
	}
	ws := rs[1].Payload
	if obj, _, err := core.DecodeRecord(ws[6:]); err != nil {
		t.Fatal(err)
	} else if b, _ := core.DecodeBytesValue(obj, codecURIEven); string(b) != "http://wikiba.se/ontology#ExternalId" {
		t.Fatalf("object %q", b)
	}
	if rs := datatypeEdges(&next, subject, "", &rep); rs != nil || next != 7 {
		t.Fatal("item: want none")
	}
	if rs := datatypeEdges(&next, subject, "nope", &rep); rs != nil || rep.Skipped["데이터 타입: nope"] != 1 {
		t.Fatal("unknown: want skipped")
	}
}
