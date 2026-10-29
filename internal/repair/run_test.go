//ff:func feature=repair type=command control=iteration dimension=1
//ff:what TestRun — a dry run reports the damage without writing; a write repairs each hash to its first occurrence, an entity range first and a line range after; a rerun resumes past the checkpoint; a checkpoint table without the key (dump_id, run), failures and a stop are returned
package repair

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	db, c := loadFixture(t)
	st, err := Run(context.Background(), db, c)
	if err != nil {
		t.Fatal(err)
	}
	if st.Spans != 3 || st.Entities != 4 || len(st.Errors) != 1 || st.Hashes != (Tally{Same: 1, Order: 1, Content: 1, Missing: 1}) || st.Replaced != 0 {
		t.Fatalf("dry run: %+v", st)
	}
	if st.Occurrences != (Tally{Same: 3, Order: 1, Content: 2, Missing: 1}) || st.ContentEntities != 2 || st.OrderEntities != 1 || st.StatementRefs != 7 || st.RefDiffs != 1 || st.RefDiffEntities[0] != "Q3" {
		t.Fatalf("dry run occurrences: %+v", st)
	}
	if got := snakList(t, db, c, "hz"); got != "" {
		t.Fatalf("dry run wrote: %q", got)
	}
	c.DryRun, c.Kind, c.Lo, c.Hi = false, "Q", 2, 3
	if st, err = Run(context.Background(), db, c); err != nil || st.Replaced != 2 || st.Last != 4 {
		t.Fatalf("entities: %+v %v", st, err)
	}
	if got := snakList(t, db, c, "hx"); got != "P123=Q7809/0 P123=Q2/1 P813=+2020-01-01T00:00:00Z/0 P248=Q36578/0" {
		t.Fatalf("hx after Q2-Q3: %q", got)
	}
	c.Kind, c.Lo, c.Hi = "", 0, 0
	if st, err = Run(context.Background(), db, c); err != nil || st.Replaced != 2 || st.Skipped != 2 || st.Hashes.Order != 2 {
		t.Fatalf("lines: %+v %v", st, err)
	}
	want := map[string]string{"hx": "P248=Q36578/0 P813=+2020-01-01T00:00:00Z/0 P123=Q7809/0 P123=Q2/1", "hy": "P854=https://example.org/a/0 P248=Q1/0", "hz": "P854=https://example.org/z/0", "hw": "P143=Q328/0"}
	for h, w := range want {
		if got := snakList(t, db, c, h); got != w {
			t.Fatalf("%s: %q want %q", h, got, w)
		}
	}
	if st, err = Run(context.Background(), db, c); err != nil || st.Planned != 0 || st.Last != 5 {
		t.Fatalf("resume: %+v %v", st, err)
	}
	c.DryRun = true
	if st, err = Run(context.Background(), db, c); err != nil || st.Hashes != (Tally{Same: 4}) || st.Occurrences.Order != 2 {
		t.Fatalf("after repair: %+v %v", st, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Run(ctx, db, c); !errors.Is(err, context.Canceled) {
		t.Fatalf("stopped: %v", err)
	}
	c.DryRun = false
	held, err := takeLock(context.Background(), db, c)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Run(context.Background(), db, c); err == nil {
		t.Fatal("lock held: want error")
	}
	releaseLock(held, c)
	if _, err := Run(ctx, db, c); err == nil {
		t.Fatal("stopped before the lock: want error")
	}
	none := c
	none.Schema += "_none"
	if _, err := Run(context.Background(), db, none); err == nil {
		t.Fatal("no schema for the work tables: want error")
	}
	db.Exec("DROP TABLE " + c.Schema + ".repair_reference_snak_checkpoint")
	db.Exec("CREATE TABLE " + c.Schema + ".repair_reference_snak_checkpoint (run text)")
	if _, err := Run(context.Background(), db, c); err == nil || !strings.Contains(err.Error(), "migrate_checkpoint_key.sql") {
		t.Fatalf("checkpoint table without the key (dump_id, run): %v", err)
	}
	none.DryRun = true
	if _, err := Run(context.Background(), db, none); err == nil {
		t.Fatal("dry run without tables: want error")
	}
	c.DryRun = true
	db.Exec("DROP TABLE " + c.Schema + ".repair_reference_snak_checkpoint")
	db.Exec("ALTER TABLE " + c.Schema + ".statement_reference RENAME TO gone")
	if _, err := Run(context.Background(), db, c); err == nil {
		t.Fatal("failed batch: want error")
	}
}
