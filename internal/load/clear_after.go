//ff:func feature=load type=writer control=iteration dimension=1
//ff:what ClearAfter — deletes the rows of dump lines after a checkpoint, which a stopped load may have committed in part
package load

import (
	"database/sql"

	"github.com/lib/pq"
)

// ClearAfter deletes, in one transaction, every row of dump id that comes from
// a line after line: batches commit in parallel, so lines past the checkpoint
// may be partly in. Rows of entity tables are found through entity.line_no and
// statement keys (statementID: dump << 55 | line << 20 | i). reference is
// shared by hash and stays (wikidata.sql 4.2 removes duplicates);
// reference_snak rows written from those lines' occurrences
// (occ_statement_id) go, so an occurrence is never in twice. qualifier,
// statement_reference and reference_snak have no line and are cleared by the
// dump's key range.
func ClearAfter(db *sql.DB, schema string, id int16, line int64) error {
	s := pq.QuoteIdentifier(schema) + "."
	after := "(SELECT entity_id FROM " + s + "entity WHERE dump_id = $1 AND line_no > $2)"
	from, to := int64(id)<<55|line<<20|(1<<20-1), int64(id)<<55|(1<<55-1)
	byLine, byKey, byDumpKey := []any{id, line}, []any{from, to}, []any{id, from}
	queries := []struct {
		q    string
		args []any
	}{
		{"DELETE FROM " + s + "label WHERE dump_id = $1 AND entity_id IN " + after, byLine},
		{"DELETE FROM " + s + "description WHERE dump_id = $1 AND entity_id IN " + after, byLine},
		{"DELETE FROM " + s + "alias WHERE dump_id = $1 AND entity_id IN " + after, byLine},
		{"DELETE FROM " + s + "sitelink WHERE dump_id = $1 AND entity_id IN " + after, byLine},
		{"DELETE FROM " + s + "qualifier WHERE statement_id > $1 AND statement_id <= $2", byKey},
		{"DELETE FROM " + s + "statement_reference WHERE statement_id > $1 AND statement_id <= $2", byKey},
		{"DELETE FROM " + s + "reference_snak WHERE occ_statement_id > $1 AND occ_statement_id <= $2", byKey},
		{"DELETE FROM " + s + "statement WHERE dump_id = $1 AND statement_id > $2", byDumpKey},
		{"DELETE FROM " + s + "entity_raw WHERE dump_id = $1 AND line_no > $2", byLine},
		{"DELETE FROM " + s + "entity WHERE dump_id = $1 AND line_no > $2", byLine},
		{"DELETE FROM " + s + "load_error WHERE dump_id = $1 AND line_no > $2", byLine},
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, q := range queries {
		if _, err := tx.Exec(q.q, q.args...); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
