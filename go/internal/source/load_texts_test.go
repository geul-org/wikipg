//ff:func feature=source type=loader control=sequence
//ff:what TestLoadTexts — Q42's ko and en labels; a closed handle and a query with the wrong columns error
package source

import (
	"testing"

	"github.com/lib/pq"
)

func TestLoadTexts(t *testing.T) {
	db := openDB(t)
	q := "SELECT language, label FROM entity_labels WHERE entity_id = $1 AND language = ANY($2) ORDER BY language"
	ts, err := loadTexts(db, q, "Q42", pq.Array([]string{"ko", "en"}))
	if err != nil || len(ts) != 2 || ts[0].Lang != "en" || ts[0].Value != "Douglas Adams" || ts[1].Lang != "ko" {
		t.Fatalf("got %+v %v", ts, err)
	}
	if _, err := loadTexts(closedDB(t), q, "Q42", pq.Array([]string{"en"})); err == nil {
		t.Fatal("closed: want error")
	}
	bad := "SELECT language, label, entity_id FROM entity_labels WHERE entity_id = $1 AND language = ANY($2)"
	if _, err := loadTexts(db, bad, "Q42", pq.Array([]string{"en"})); err == nil {
		t.Fatal("three columns: want scan error")
	}
}
