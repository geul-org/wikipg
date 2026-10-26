//ff:func feature=source type=loader control=sequence
//ff:what TestEntityHead — Q42 exists with no datatype, P31 with wikibase-item, a never-created ID does not, a closed handle errors
package source

import "testing"

func TestEntityHead(t *testing.T) {
	db, s := fixtureDB(t)
	q := queriesFor(s).Exists
	if dt, ok, err := entityHead(db, q, "Q42"); err != nil || !ok || dt != "" {
		t.Fatalf("Q42: %q %v %v", dt, ok, err)
	}
	if dt, ok, err := entityHead(db, q, "P31"); err != nil || !ok || dt != "wikibase-item" {
		t.Fatalf("P31: %q %v %v", dt, ok, err)
	}
	if _, ok, err := entityHead(db, q, "Q6"); err != nil || ok {
		t.Fatalf("Q6: %v %v", ok, err)
	}
	if _, _, err := entityHead(closedDB(t), q, "Q42"); err == nil {
		t.Fatal("closed: want error")
	}
}
