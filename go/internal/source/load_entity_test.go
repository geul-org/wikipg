//ff:func feature=source type=loader control=sequence
//ff:what TestLoadEntity — Q42 with labels, aliases and claims; a never-created ID is false; a closed handle errors
package source

import "testing"

func TestLoadEntity(t *testing.T) {
	db := openDB(t)
	e, ok, err := LoadEntity(db, "Q42", []string{"ko", "en"})
	if err != nil || !ok || e.ID != "Q42" || len(e.Labels) != 2 || len(e.Aliases) == 0 || len(e.Claims) < 100 {
		t.Fatalf("Q42: %v %v %+v", ok, err, e.Labels)
	}
	if _, ok, err := LoadEntity(db, "Q0", []string{"en"}); err != nil || ok {
		t.Fatalf("Q0: %v %v", ok, err)
	}
	if _, _, err := LoadEntity(closedDB(t), "Q42", []string{"en"}); err == nil {
		t.Fatal("closed: want error")
	}
}
