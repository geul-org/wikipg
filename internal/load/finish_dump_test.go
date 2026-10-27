//ff:func feature=load type=writer control=sequence
//ff:what TestFinishDump — load_finished is set; a closed handle fails
package load

import "testing"

func TestFinishDump(t *testing.T) {
	db, schema := testSchema(t)
	InsertDump(db, schema, sampleDump())
	if err := FinishDump(db, schema, 1); err != nil {
		t.Fatal(err)
	}
	var done bool
	db.QueryRow("SELECT load_finished IS NOT NULL FROM " + schema + ".dump").Scan(&done)
	if !done {
		t.Fatal("not finished")
	}
	if err := FinishDump(closedDB(t), schema, 1); err == nil {
		t.Fatal("closed: want error")
	}
}
