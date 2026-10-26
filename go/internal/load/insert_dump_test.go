//ff:func feature=load type=writer control=sequence
//ff:what TestInsertDump — the dump row is written once with load_started; a second insert keeps the first; no published time and no MD5 are NULL; a closed handle fails
package load

import (
	"testing"
	"time"
)

func TestInsertDump(t *testing.T) {
	db, schema := testSchema(t)
	d := sampleDump()
	if err := InsertDump(db, schema, d); err != nil {
		t.Fatal(err)
	}
	d.FileName = "other"
	if err := InsertDump(db, schema, d); err != nil {
		t.Fatal(err)
	}
	var name string
	var started bool
	db.QueryRow("SELECT file_name, load_started IS NOT NULL FROM "+schema+".dump WHERE dump_id = 1").Scan(&name, &started)
	if name != "wikidata-20261019-all.json.bz2" || !started {
		t.Fatalf("%q %v", name, started)
	}
	d2 := sampleDump()
	d2.ID, d2.PublishedAt, d2.MD5 = 2, time.Time{}, ""
	if err := InsertDump(db, schema, d2); err != nil {
		t.Fatal(err)
	}
	var noPub, noMD5 bool
	db.QueryRow("SELECT published_at IS NULL, md5 IS NULL FROM "+schema+".dump WHERE dump_id = 2").Scan(&noPub, &noMD5)
	if !noPub || !noMD5 {
		t.Fatalf("NULLs: %v %v", noPub, noMD5)
	}
	if err := InsertDump(closedDB(t), schema, d); err == nil {
		t.Fatal("closed: want error")
	}
}
