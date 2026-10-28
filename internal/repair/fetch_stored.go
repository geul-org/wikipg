//ff:func feature=repair type=reader control=iteration dimension=1
//ff:what fetchStored — the stored reference_snak rows of some hashes in ordinal order, by hash, the value in jsonb text form; one primary-key lookup per hash
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// fetchStored returns the reference_snak rows of hashes in the dump, by hash
// in ordinal order, the value as PostgreSQL prints the jsonb. Each hash is
// one lookup on the primary key: the LATERAL subquery with OFFSET 0 keeps the
// planner from a parallel sequential scan of the whole table, which it
// chose for hash = ANY(array) with tens of thousands of hashes (row estimate
// in the millions).
func fetchStored(db *sql.DB, c Config, hashes []string) (map[string][]Snak, error) {
	rows, err := db.Query("SELECT r.hash, r.ordinal, r.property, r.value_ordinal, r.snaktype::text, r.datatype, r.value_type, r.value::text, r.value_entity, r.value_json "+
		"FROM unnest($2::text[]) AS h(hash) CROSS JOIN LATERAL (SELECT * FROM "+pq.QuoteIdentifier(c.Schema)+".reference_snak "+
		"WHERE hash = h.hash AND dump_id = $1 OFFSET 0) r ORDER BY r.hash, r.ordinal", c.DumpID, pq.Array(hashes))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[string][]Snak{}
	for rows.Next() {
		var h string
		var s Snak
		if err := rows.Scan(&h, &s.Ordinal, &s.Property, &s.ValueOrdinal, &s.SnakType, &s.Datatype, &s.ValueType, &s.Value, &s.ValueEntity, &s.ValueJSON); err != nil {
			return nil, err
		}
		out[h] = append(out[h], s)
	}
	return out, rows.Err()
}
