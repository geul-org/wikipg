//ff:func feature=source type=loader control=sequence
//ff:what TestLoadEntity — Q42 with labels, descriptions, aliases and claims of its dump only; a missing ID is false; a bad schema and a closed handle error
package source

import "testing"

func TestLoadEntity(t *testing.T) {
	db, s := fixtureDB(t)
	e, ok, err := LoadEntity(db, s, "Q42", []string{"ko", "en"})
	if err != nil || !ok || e.ID != "Q42" || len(e.Labels) != 2 || len(e.Descriptions) != 2 || len(e.Aliases) != 2 || len(e.Claims) != 6 || len(e.Claims[0].Qualifiers) != 3 {
		t.Fatalf("Q42: %v %v %+v", ok, err, e)
	}
	if _, ok, err := LoadEntity(db, s, "Q6", []string{"en"}); err != nil || ok {
		t.Fatalf("Q6: %v %v", ok, err)
	}
	if _, ok, err := LoadEntity(db, s, "Q99", []string{"en"}); err != nil || ok {
		t.Fatalf("Q99 of another dump: %v %v", ok, err)
	}
	if _, _, err := LoadEntity(db, Store{Schema: "x;drop", Dump: 1}, "Q42", []string{"en"}); err == nil {
		t.Fatal("bad schema: want error")
	}
	if _, _, err := LoadEntity(closedDB(t), s, "Q42", []string{"en"}); err == nil {
		t.Fatal("closed: want error")
	}
}
