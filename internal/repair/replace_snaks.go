//ff:func feature=repair type=writer control=sequence
//ff:what replaceSnaks — in a transaction, deletes the stored reference_snak rows of some hashes and inserts their occurrence's rows
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// replaceSnaks deletes the reference_snak rows of the hashes of checks in tx
// (found by one primary-key lookup per hash, deleted by row address, so the
// planner cannot turn the delete into a scan of the table) and inserts each one's Rows in one INSERT … SELECT FROM unnest (the columns
// without the loader's staging columns, so it works on a table that went
// through wikidata.sql 4.2). The table's indexes (reference_snak_pkey,
// reference_snak_value_entity) follow the rows. Nothing for no checks.
func replaceSnaks(tx *sql.Tx, c Config, checks []HashCheck) error {
	if len(checks) == 0 {
		return nil
	}
	s := pq.QuoteIdentifier(c.Schema)
	if _, err := tx.Exec("DELETE FROM "+s+".reference_snak WHERE ctid = ANY(ARRAY(SELECT r.ctid FROM unnest($2::text[]) AS h(hash) CROSS JOIN LATERAL "+
		"(SELECT ctid FROM "+s+".reference_snak WHERE hash = h.hash AND dump_id = $1 OFFSET 0) r))", c.DumpID, pq.Array(hashesOf(checks))); err != nil {
		return err
	}
	k := columnsOf(checks)
	_, err := tx.Exec("INSERT INTO "+s+".reference_snak (hash, dump_id, ordinal, property, value_ordinal, snaktype, datatype, value_type, value, value_entity, value_json) "+
		"SELECT h, d, o, p, vo, st::"+s+".snak_type, dt, vt, v::jsonb, ve, vj FROM unnest($1::text[], $2::smallint[], $3::smallint[], $4::text[], $5::smallint[], $6::text[], "+
		"$7::text[], $8::text[], $9::text[], $10::text[], $11::text[]) AS t(h, d, o, p, vo, st, dt, vt, v, ve, vj)",
		pq.Array(k.Hash), pq.Array(k.DumpID), pq.Array(k.Ordinal), pq.Array(k.Property), pq.Array(k.ValueOrdinal), pq.Array(k.SnakType),
		pq.Array(k.Datatype), pq.Array(k.ValueType), pq.Array(k.Value), pq.Array(k.ValueEntity), pq.Array(k.ValueJSON))
	return err
}
