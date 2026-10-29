//ff:func feature=repair type=writer control=sequence
//ff:what TestEnsureTables — creates both work tables of a repair, or the order table and its checkpoint for an order build, in the schema; a second time is a no-op; a missing schema fails
package repair

import "testing"

func TestEnsureTables(t *testing.T) {
	db, schema := testSchema(t)
	c := Config{Schema: schema}
	if err := ensureTables(db, c); err != nil {
		t.Fatal(err)
	}
	if err := ensureTables(db, c); err != nil {
		t.Fatal("again:", err)
	}
	var n int
	db.QueryRow("SELECT count(*) FROM information_schema.tables WHERE table_schema = $1 AND table_name LIKE 'repair_reference_snak_%'", schema).Scan(&n)
	if n != 2 {
		t.Fatalf("tables %d", n)
	}
	if err := ensureTables(db, Config{Schema: schema, Order: true}); err != nil {
		t.Fatal("order:", err)
	}
	db.QueryRow("SELECT count(*) FROM information_schema.tables WHERE table_schema = $1 AND table_name LIKE 'reference_snak_order%'", schema).Scan(&n)
	if n != 2 {
		t.Fatalf("order tables %d", n)
	}
	if err := ensureTables(db, Config{Schema: schema + "_none"}); err == nil {
		t.Fatal("missing schema: want error")
	}
}
