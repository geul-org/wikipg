//ff:func feature=source type=loader control=sequence
//ff:what TestEntityExists — Q42 exists, a never-created ID does not, a closed handle errors
package source

import "testing"

func TestEntityExists(t *testing.T) {
	db := openDB(t)
	if ok, err := entityExists(db, defaultQueries().Exists, "Q42"); err != nil || !ok {
		t.Fatalf("Q42: %v %v", ok, err)
	}
	if ok, err := entityExists(db, defaultQueries().Exists, "Q0"); err != nil || ok {
		t.Fatalf("Q0: %v %v", ok, err)
	}
	if _, err := entityExists(closedDB(t), defaultQueries().Exists, "Q42"); err == nil {
		t.Fatal("closed: want error")
	}
}
