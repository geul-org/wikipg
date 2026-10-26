//ff:func feature=load type=writer control=sequence
//ff:what TestSaveCheckpoint — the line and counts are stored on the dump row; a closed handle fails
package load

import "testing"

func TestSaveCheckpoint(t *testing.T) {
	db, schema := testSchema(t)
	InsertDump(db, schema, sampleDump())
	if err := SaveCheckpoint(db, schema, 1, 10, 9, 0); err != nil {
		t.Fatal(err)
	}
	var l int64
	db.QueryRow("SELECT checkpoint_line FROM " + schema + ".dump").Scan(&l)
	if l != 10 {
		t.Fatalf("line %d", l)
	}
	if err := SaveCheckpoint(closedDB(t), schema, 1, 10, 9, 0); err == nil {
		t.Fatal("closed: want error")
	}
}
