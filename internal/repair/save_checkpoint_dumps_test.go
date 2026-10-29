//ff:func feature=repair type=writer control=sequence
//ff:what TestSaveCheckpointDumps — regression: dumps 1 and 2 saving and resuming the same range name (lines:0-0) in turn keep their own lines, in both modes
package repair

import (
	"io"
	"testing"
)

// Before 2026-10-29 the checkpoint key was run alone and saving was
// ON CONFLICT (run): dump 2 under lines:0-0 moved dump 1's line, and dump 2
// itself read nothing back.
func TestSaveCheckpointDumps(t *testing.T) {
	db, schema := testSchema(t)
	d1 := Config{Schema: schema, DumpID: 1, Log: io.Discard}
	d2 := d1
	d2.DumpID = 2
	o1, o2 := d1, d2
	o1.Order, o2.Order = true, true
	if d1.Run() != "lines:0-0" || d2.Run() != d1.Run() {
		t.Fatalf("range names %s %s", d1.Run(), d2.Run())
	}
	ensureTables(db, d1)
	ensureTables(db, o1)
	save := func(c Config, line int64) {
		t.Helper()
		tx, _ := db.Begin()
		if err := saveCheckpoint(tx, c, line); err != nil {
			t.Fatal(err)
		}
		tx.Commit()
	}
	resume := func(c Config) int64 {
		t.Helper()
		n, err := resumeLine(db, c)
		if err != nil {
			t.Fatal(err)
		}
		return n
	}
	save(d1, 121971282)
	if n := resume(d2); n != 0 {
		t.Fatalf("dump 2 before its first batch resumes after dump 1's line: %d", n)
	}
	save(d2, 1000)
	if a, b := resume(d1), resume(d2); a != 121971282 || b != 1000 {
		t.Fatalf("after dump 2 saved: dump 1 %d, dump 2 %d", a, b)
	}
	save(d1, 121971283)
	save(d2, 1707633)
	if a, b := resume(d1), resume(d2); a != 121971283 || b != 1707633 {
		t.Fatalf("in turn: dump 1 %d, dump 2 %d", a, b)
	}
	save(o2, 50)
	save(o1, 70)
	save(o2, 60)
	if a, b := resume(o1), resume(o2); a != 70 || b != 60 {
		t.Fatalf("order build: dump 1 %d, dump 2 %d", a, b)
	}
	if a, b := resume(d1), resume(d2); a != 121971283 || b != 1707633 {
		t.Fatalf("repair moved by the order build: dump 1 %d, dump 2 %d", a, b)
	}
	var n int
	db.QueryRow("SELECT (SELECT count(*) FROM " + schema + ".repair_reference_snak_checkpoint WHERE run = 'lines:0-0') + " +
		"(SELECT count(*) FROM " + schema + ".reference_snak_order_checkpoint WHERE run = 'lines:0-0')").Scan(&n)
	if n != 4 {
		t.Fatalf("rows %d, want one per dump and mode", n)
	}
}
