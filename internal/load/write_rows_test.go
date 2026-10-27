//ff:func feature=load type=writer control=sequence
//ff:what TestWriteRows — the sample item lands in every table with jsonb values and arrays intact; a closed handle and a missing schema fail; a lexeme lands in its tables
package load

import (
	"testing"
)

func TestWriteRows(t *testing.T) {
	db, schema := testSchema(t)
	if err := InsertDump(db, schema, sampleDump()); err != nil {
		t.Fatal(err)
	}
	r := NewRows()
	r.AddEntity(Config{DumpID: 1, Raw: true}, 5, []byte(`{"id":"Q42"}`), sampleEntity(), NewRefSet(64))
	r.AddEntity(Config{DumpID: 1, Raw: true}, 7, []byte(`{"id":"L4"}`), sampleLexeme(), NewRefSet(64))
	r.Error(1, 6, "", "parse", "x", []byte("y"))
	if err := WriteRows(db, schema, r); err != nil {
		t.Fatal(err)
	}
	var unit, badge, snak string
	var quals int
	err := db.QueryRow("SELECT s.value->>'unit', (SELECT badges[1] FROM "+schema+".sitelink), (SELECT snaktype::text FROM "+schema+".statement WHERE property = 'P40'), "+
		"(SELECT count(*) FROM "+schema+".qualifier) FROM "+schema+".statement s WHERE s.property = 'P2048'").Scan(&unit, &badge, &snak, &quals)
	if err != nil || unit != "http://www.wikidata.org/entity/Q11573" || badge != "Q17437796" || snak != "novalue" || quals != 2 {
		t.Fatalf("%q %q %q %d %v", unit, badge, snak, quals, err)
	}
	var feature, gloss string
	var forms, entities int
	err = db.QueryRow("SELECT (SELECT grammatical_features[1] FROM "+schema+".form), (SELECT value FROM "+schema+".sense_gloss WHERE language = 'en'), "+
		"(SELECT count(*) FROM "+schema+".form_representation), (SELECT count(*) FROM "+schema+".entity WHERE kind = 'L')").Scan(&feature, &gloss, &forms, &entities)
	if err != nil || feature != "Q10345583" || gloss != "to sail on a sailboard" || forms != 1 || entities != 3 {
		t.Fatalf("lexeme: %q %q %d %d %v", feature, gloss, forms, entities, err)
	}
	if err := WriteRows(closedDB(t), schema, r); err == nil {
		t.Fatal("closed: want error")
	}
	if err := WriteRows(db, schema+"_none", r); err == nil {
		t.Fatal("missing schema: want error")
	}
}
