//ff:func feature=source type=loader control=sequence
//ff:what TestLoadClaims — Q42's triples in ID order; a closed handle and a query with the wrong columns error
package source

import "testing"

func TestLoadClaims(t *testing.T) {
	db := openDB(t)
	q := "SELECT id, property, object_type, coalesce(object_value, ''), coalesce(rank, '') FROM triples WHERE subject = $1 ORDER BY id"
	cs, err := loadClaims(db, q, "Q42")
	if err != nil || len(cs) < 100 || cs[0].ID >= cs[1].ID || cs[0].Property == "" || cs[0].Rank == "" {
		t.Fatalf("got %d claims %v", len(cs), err)
	}
	if _, err := loadClaims(closedDB(t), q, "Q42"); err == nil {
		t.Fatal("closed: want error")
	}
	if _, err := loadClaims(db, "SELECT id FROM triples WHERE subject = $1", "Q42"); err == nil {
		t.Fatal("one column: want scan error")
	}
}
