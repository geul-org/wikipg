//ff:func feature=load type=writer control=sequence
//ff:what TestReadCheckpoint — a new dump reads zero; a saved checkpoint reads back; an unknown dump fails
package load

import "testing"

func TestReadCheckpoint(t *testing.T) {
	db, schema := testSchema(t)
	InsertDump(db, schema, sampleDump())
	if l, e, x, err := ReadCheckpoint(db, schema, 1); err != nil || l != 0 || e != 0 || x != 0 {
		t.Fatalf("new: %d %d %d %v", l, e, x, err)
	}
	SaveCheckpoint(db, schema, 1, 99, 98, 1)
	if l, e, x, err := ReadCheckpoint(db, schema, 1); err != nil || l != 99 || e != 98 || x != 1 {
		t.Fatalf("saved: %d %d %d %v", l, e, x, err)
	}
	if _, _, _, err := ReadCheckpoint(db, schema, 9); err == nil {
		t.Fatal("unknown dump: want error")
	}
}
