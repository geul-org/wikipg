//ff:func feature=source type=loader control=sequence
//ff:what TestEntityExists — Q42 exists in its dump, a never-created ID does not, a closed handle errors
package source

import "testing"

func TestEntityExists(t *testing.T) {
	db, s := fixtureDB(t)
	q := queriesFor(s).Exists
	if ok, err := entityExists(db, q, "Q42"); err != nil || !ok {
		t.Fatalf("Q42: %v %v", ok, err)
	}
	if ok, err := entityExists(db, q, "Q6"); err != nil || ok {
		t.Fatalf("Q6: %v %v", ok, err)
	}
	if _, err := entityExists(closedDB(t), q, "Q42"); err == nil {
		t.Fatal("closed: want error")
	}
}
