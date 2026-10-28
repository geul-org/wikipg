//ff:func feature=repair type=writer control=iteration dimension=1
//ff:what markDone — records in a transaction the hashes processed with their occurrence and verdict (an earlier occurrence overwrites)
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// markDone records each of checks in repair_reference_snak_done in tx: the
// hash, the occurrence it was processed at and its verdict. A hash already
// there (processed later in the dump by an earlier run of another range)
// takes the new occurrence. Nothing for no checks.
func markDone(tx *sql.Tx, c Config, checks []HashCheck) error {
	if len(checks) == 0 {
		return nil
	}
	var hashes, verdicts []string
	var sids, ords []int64
	for _, h := range checks {
		hashes, verdicts = append(hashes, h.Hash), append(verdicts, h.Verdict.String())
		sids, ords = append(sids, h.At.Statement), append(ords, int64(h.At.Ordinal))
	}
	_, err := tx.Exec("INSERT INTO "+pq.QuoteIdentifier(c.Schema)+".repair_reference_snak_done (hash, dump_id, occ_statement_id, occ_ordinal, verdict) "+
		"SELECT h, $1, s, o, v FROM unnest($2::text[], $3::bigint[], $4::smallint[], $5::text[]) AS t(h, s, o, v) "+
		"ON CONFLICT (hash, dump_id) DO UPDATE SET occ_statement_id = EXCLUDED.occ_statement_id, occ_ordinal = EXCLUDED.occ_ordinal, verdict = EXCLUDED.verdict",
		c.DumpID, pq.Array(hashes), pq.Array(sids), pq.Array(ords), pq.Array(verdicts))
	return err
}
