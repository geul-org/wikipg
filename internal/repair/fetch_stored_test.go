//ff:func feature=repair type=reader control=sequence
//ff:what TestFetchStored — stored rows by hash in ordinal order with jsonb text values; a hash without rows is absent; a NULL property and a missing table fail
package repair

import "testing"

func TestFetchStored(t *testing.T) {
	db, c := loadFixture(t)
	got, err := fetchStored(db, c, []string{"hw", "hy", "hz"})
	if err != nil || len(got) != 2 || len(got["hy"]) != 2 || got["hy"][0].Property != "P248" || got["hw"][0].Value.String != `{"id": "Q328", "numeric-id": 328, "entity-type": "item"}` {
		t.Fatalf("%+v %v", got, err)
	}
	db.Exec("ALTER TABLE " + c.Schema + ".reference_snak ALTER property DROP NOT NULL")
	db.Exec("INSERT INTO " + c.Schema + ".reference_snak (hash, dump_id, ordinal, property, value_ordinal, snaktype) VALUES ('hn', 1, 0, NULL, 0, 'novalue')")
	if _, err := fetchStored(db, c, []string{"hn"}); err == nil {
		t.Fatal("NULL property: want error")
	}
	if _, err := fetchStored(db, Config{Schema: c.Schema + "_none"}, []string{"hw"}); err == nil {
		t.Fatal("missing table: want error")
	}
}
