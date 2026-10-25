//ff:func feature=publish type=encoder control=sequence
//ff:what TestNameIndex — sorted, deduplicated BINDs that resolve.ParseDocumentNames reads back; a bad Issuer fails
package publish

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
)

func TestNameIndex(t *testing.T) {
	b, err := NameIndex([]core.Word{1, 0}, "q", []uint64{42, 1, 42})
	if err != nil {
		t.Fatal(err)
	}
	names, err := resolve.ParseDocumentNames(b)
	if err != nil || len(names) != 2 || names["q42"] != (resolve.DocKey{Issuer: "0x00010000", Doc: "0x000000000000002a"}) || names["q1"].Doc != "0x0000000000000001" {
		t.Fatalf("got %v %v", names, err)
	}
	if _, err := NameIndex([]core.Word{0, 1}, "q", []uint64{1}); err == nil {
		t.Fatal("bad Issuer: want error")
	}
}
