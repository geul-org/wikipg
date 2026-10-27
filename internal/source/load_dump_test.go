//ff:func feature=source type=loader control=sequence
//ff:what TestLoadDump — the fixture dump's URL, download time, loader and commit; a missing dump and a bad schema error
package source

import (
	"strings"
	"testing"
)

func TestLoadDump(t *testing.T) {
	db, s := fixtureDB(t)
	d, err := LoadDump(db, s)
	if err != nil || !strings.HasSuffix(d.URL, "wikidata-20261019-all.json.bz2") || !strings.HasPrefix(d.SHA256, "a249a19d") || d.DownloadedAt.Unix() != 1792996483 || d.Loader != "github.com/geul-org/wikiport/cmd/wikiload" || len(d.LoaderCommit) != 40 {
		t.Fatalf("got %+v %v", d, err)
	}
	if _, err := LoadDump(db, Store{Schema: s.Schema, Dump: 9}); err == nil {
		t.Fatal("missing dump: want error")
	}
	if _, err := LoadDump(db, Store{Schema: "Bad", Dump: 1}); err == nil {
		t.Fatal("bad schema: want error")
	}
}
