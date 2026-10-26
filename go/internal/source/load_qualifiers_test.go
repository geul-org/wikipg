//ff:func feature=source type=loader control=iteration dimension=1
//ff:what TestLoadQualifiers — Q42's qualifiers include St John's College start 1971 (triple 1756); a closed handle and wrong columns error
package source

import (
	"strings"
	"testing"
)

func TestLoadQualifiers(t *testing.T) {
	db := openDB(t)
	q := defaultQueries().Qualifiers
	qs, err := loadQualifiers(db, q, "Q42")
	if err != nil || len(qs) == 0 {
		t.Fatalf("%d %v", len(qs), err)
	}
	found := false
	for _, x := range qs {
		if x.TripleID == 1756 && x.Property == "P580" && x.Type == "time" && strings.Contains(x.Value, "+1971-") {
			found = true
		}
	}
	if !found {
		t.Fatal("triple 1756 P580 1971 not found")
	}
	if _, err := loadQualifiers(closedDB(t), q, "Q42"); err == nil {
		t.Fatal("closed: want error")
	}
	if _, err := loadQualifiers(db, "SELECT 1, 2 FROM triples WHERE subject = $1", "Q42"); err == nil {
		t.Fatal("columns: want error")
	}
}
