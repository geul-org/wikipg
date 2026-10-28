//ff:func feature=repair type=reader control=iteration dimension=1
//ff:what TestCheckSpan — a batch's hashes, occurrences and statement_reference difference against the damaged fixture; each read failing fails it
package repair

import "testing"

func TestCheckSpan(t *testing.T) {
	db, c := loadFixture(t)
	f, err := checkSpan(db, c, Span{Seq: 4, From: 2, To: 5, Last: 5})
	if err != nil || f.Seq != 4 || f.Last != 5 || f.Entities != 4 || len(f.Hashes) != 4 || f.Hashes[0].Hash != "hx" || f.Hashes[0].Verdict != Content {
		t.Fatalf("%+v %v", f, err)
	}
	if f.StatementRefs != 7 || f.RefDiffs != 1 || len(f.RefDiffEntities) != 1 || f.RefDiffEntities[0] != "Q3" || len(f.Errors) != 1 {
		t.Fatalf("refs %+v", f)
	}
	bad := c
	bad.Schema = c.Schema + "_none"
	if _, err := checkSpan(db, bad, Span{From: 2, To: 2}); err == nil {
		t.Fatal("missing schema: want error")
	}
	// A value jsonb rejects (a lone surrogate) fails the normalization.
	db.Exec("INSERT INTO " + c.Schema + `.entity_raw VALUES ('Q8', 1, 8, '{"id":"Q8","claims":{"P1":[{"id":"Q8$1","rank":"normal","mainsnak":{"snaktype":"novalue","property":"P1"},` +
		`"references":[{"hash":"hs","snaks":{"P1":[{"snaktype":"value","property":"P1","datavalue":{"value":"\ud800","type":"string"}}]},"snaks-order":["P1"]}]}]}}')`)
	if _, err := checkSpan(db, c, Span{From: 8, To: 8}); err == nil {
		t.Fatal("lone surrogate: want error")
	}
	for _, table := range []string{"reference_snak", "statement_reference"} {
		db.Exec("ALTER TABLE " + c.Schema + "." + table + " RENAME TO " + table + "_gone")
		if _, err := checkSpan(db, c, Span{From: 2, To: 2}); err == nil {
			t.Fatalf("no %s: want error", table)
		}
	}
}
