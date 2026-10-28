//ff:func feature=repair type=reader control=sequence
//ff:what TestReadSpan — a window and a line list read in line order with occurrences and statement_reference rows; a broken line is an error entry; a NULL raw and a missing table fail
package repair

import "testing"

func TestReadSpan(t *testing.T) {
	db, c := loadFixture(t)
	d, err := readSpan(db, c, Span{From: 1, To: 5})
	if err != nil || len(d.Lines) != 4 || d.IDs[4] != "Q3" || len(d.Occs) != 7 || len(d.Refs) != 7 || len(d.Errors) != 1 || d.Bytes == 0 {
		t.Fatalf("window: %+v %v", d, err)
	}
	if d, err := readSpan(db, c, Span{Lines: []int64{4, 2}}); err != nil || len(d.Lines) != 2 || d.Lines[0] != 2 || d.IDs[2] != "Q1" {
		t.Fatalf("list: %+v %v", d, err)
	}
	db.Exec("ALTER TABLE " + c.Schema + ".entity_raw ALTER raw DROP NOT NULL")
	db.Exec("INSERT INTO " + c.Schema + ".entity_raw VALUES ('Q9', 1, 9, NULL)")
	if _, err := readSpan(db, c, Span{From: 9, To: 9}); err == nil {
		t.Fatal("NULL raw: want error")
	}
	if _, err := readSpan(db, Config{Schema: c.Schema + "_none"}, Span{From: 1, To: 2}); err == nil {
		t.Fatal("missing table: want error")
	}
}
