//ff:func feature=load type=model control=sequence
//ff:what TestRowsAddSitelinks — one row per site with badges; no badges is an empty array; a NUL title is cleaned with an error
package load

import (
	"database/sql/driver"
	"testing"

	"github.com/geul-org/wikiport/go/internal/dump"
)

func TestRowsAddSitelinks(t *testing.T) {
	r := NewRows()
	e := dump.Entity{ID: "Q1", Sitelinks: map[string]dump.Sitelink{"enwiki": {Site: "enwiki", Title: "A\x00"}}}
	r.addSitelinks(Config{DumpID: 1}, 3, e)
	row := r.byTable["sitelink"][0]
	if r.Count("sitelink") != 1 || row[3] != "A" || r.Errors != 1 {
		t.Fatalf("got %v", row)
	}
	if v, err := row[4].(driver.Valuer).Value(); err != nil || v != "{}" {
		t.Fatalf("badges %v %v", v, err)
	}
}
