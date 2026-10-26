//ff:func feature=source type=loader control=iteration dimension=1
//ff:what TestLoadEntityWith — Q42 loads; a missing ID is false; a failing existence, label, description, alias, claim or qualifier query errors
package source

import "testing"

func TestLoadEntityWith(t *testing.T) {
	db := openDB(t)
	q := defaultQueries()
	if e, ok, err := loadEntityWith(db, q, "Q42", []string{"en"}); err != nil || !ok || len(e.Labels) != 1 || len(e.Descriptions) != 1 {
		t.Fatalf("Q42: %v %v", ok, err)
	}
	if _, ok, err := loadEntityWith(db, q, "Q0", []string{"en"}); ok || err != nil {
		t.Fatalf("Q0: %v %v", ok, err)
	}
	bad := "SELECT nope FROM nowhere"
	for i, b := range []entityQueries{
		{Exists: bad, Labels: q.Labels, Descriptions: q.Descriptions, Aliases: q.Aliases, Claims: q.Claims, Qualifiers: q.Qualifiers},
		{Exists: q.Exists, Labels: bad, Descriptions: q.Descriptions, Aliases: q.Aliases, Claims: q.Claims, Qualifiers: q.Qualifiers},
		{Exists: q.Exists, Labels: q.Labels, Descriptions: bad, Aliases: q.Aliases, Claims: q.Claims, Qualifiers: q.Qualifiers},
		{Exists: q.Exists, Labels: q.Labels, Descriptions: q.Descriptions, Aliases: bad, Claims: q.Claims, Qualifiers: q.Qualifiers},
		{Exists: q.Exists, Labels: q.Labels, Descriptions: q.Descriptions, Aliases: q.Aliases, Claims: bad, Qualifiers: q.Qualifiers},
		{Exists: q.Exists, Labels: q.Labels, Descriptions: q.Descriptions, Aliases: q.Aliases, Claims: q.Claims, Qualifiers: bad},
	} {
		if _, _, err := loadEntityWith(db, b, "Q42", []string{"en"}); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
