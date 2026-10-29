//ff:func feature=repair type=reader control=sequence
//ff:what TestResumeLine — 0 without a checkpoint table; a writing run resumes after its dump's line and says so, a dry run reports it and starts at 0; an old key and an unreadable table fail
package repair

import (
	"strings"
	"testing"
)

func TestResumeLine(t *testing.T) {
	db, schema := testSchema(t)
	var log strings.Builder
	c := Config{Schema: schema, DumpID: 1, Log: &log}
	if n, err := resumeLine(db, c); err != nil || n != 0 || log.Len() != 0 {
		t.Fatalf("no table: %d %v %q", n, err, log.String())
	}
	ensureTables(db, c)
	tx, _ := db.Begin()
	saveCheckpoint(tx, c, 8)
	tx.Commit()
	if n, err := resumeLine(db, c); err != nil || n != 8 || log.String() != "checkpoint lines:0-0 dump 1: line 8 (resuming after it)\n" {
		t.Fatalf("write: %d %v %q", n, err, log.String())
	}
	log.Reset()
	c.DryRun = true
	if n, err := resumeLine(db, c); err != nil || n != 0 || !strings.Contains(log.String(), "line 8 (dry run") {
		t.Fatalf("dry run: %d %v %q", n, err, log.String())
	}
	log.Reset()
	c.DumpID, c.DryRun = 2, false
	if n, err := resumeLine(db, c); err != nil || n != 0 || !strings.Contains(log.String(), "dump 2: line 0") {
		t.Fatalf("other dump: %d %v %q", n, err, log.String())
	}
	s := schema + "."
	db.Exec("DROP TABLE " + s + "repair_reference_snak_checkpoint")
	db.Exec("CREATE TABLE " + s + "repair_reference_snak_checkpoint (run text PRIMARY KEY, dump_id smallint NOT NULL, line_no bigint NOT NULL)")
	if _, err := resumeLine(db, c); err == nil || !strings.Contains(err.Error(), "migrate_checkpoint_key.sql") {
		t.Fatalf("old key: %v", err)
	}
	db.Exec("DROP TABLE " + s + "repair_reference_snak_checkpoint")
	db.Exec("CREATE TABLE " + s + "repair_reference_snak_checkpoint (dump_id smallint, run text, PRIMARY KEY (dump_id, run))")
	if _, err := resumeLine(db, c); err == nil {
		t.Fatal("no line_no column: want error")
	}
}
