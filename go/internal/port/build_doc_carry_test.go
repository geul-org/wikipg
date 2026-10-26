//ff:func feature=port type=encoder control=sequence
//ff:what TestBuildDocCarriesIDs — rebuilding with one claim changed keeps every other Edge Id and gives the changed claim the next ID, never reusing the old one
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/delta"
	"github.com/geul-org/geul/go/pkg/inspect"

	"github.com/geul-org/wikiport/go/internal/source"
)

func TestBuildDocCarriesIDs(t *testing.T) {
	claim := func(v string) source.Claim {
		return source.Claim{Property: "P31", Type: "wikibase-entityid", Value: v, Rank: "normal"}
	}
	e := source.Entity{ID: "Q42", Labels: []source.Text{{Lang: "en", Value: "Douglas Adams"}}, Claims: []source.Claim{claim("Q5"), claim("Q36180")}}
	v1, _, err := BuildDoc(e, []core.Word{1, 0}, prov(), freshIDs(1))
	if err != nil {
		t.Fatal(err)
	}
	e.Claims[1] = claim("Q482980")
	ids, _ := delta.NewLocalIDs(v1, 0)
	v2, _, err := BuildDoc(e, []core.Word{1, 0}, prov(), ids)
	if err != nil {
		t.Fatal(err)
	}
	d1, _ := inspect.Decode(v1)
	d2, _ := inspect.Decode(v2)
	if d1.Body[0].ID != d2.Body[0].ID || d1.Body[1].ID != d2.Body[1].ID || d2.Body[2].ID != 4 || d1.Body[2].ID != 3 {
		t.Fatalf("v1 %d %d %d, v2 %d %d %d", d1.Body[0].ID, d1.Body[1].ID, d1.Body[2].ID, d2.Body[0].ID, d2.Body[1].ID, d2.Body[2].ID)
	}
	if gone := ids.Unused(); len(gone) != 1 || gone[0] != 3 {
		t.Fatalf("removed %v", gone)
	}
}
