//ff:func feature=repair type=reader control=iteration dimension=1
//ff:what fetchStatementRefs — the stored statement_reference rows of the statements of some dump lines, one primary-key range scan per line
package repair

import (
	"database/sql"

	"github.com/lib/pq"
)

// fetchStatementRefs returns the statement_reference rows whose statements
// are on lines, in key order. Keys are dump << 55 | line << 20 | i, so each
// line is one range of the primary key; the LATERAL subquery with OFFSET 0
// keeps the planner from turning the ranges into one scan of the table
// (otherwise it chose a parallel sequential scan of all 1.3 billion rows).
func fetchStatementRefs(db *sql.DB, c Config, lines []int64) ([]stmtRef, error) {
	rows, err := db.Query("SELECT r.statement_id, r.ordinal, r.hash FROM unnest($2::bigint[]) AS l(line) CROSS JOIN LATERAL "+
		"(SELECT statement_id, ordinal, hash FROM "+pq.QuoteIdentifier(c.Schema)+".statement_reference "+
		"WHERE statement_id BETWEEN (($1::bigint << 55) | (l.line << 20)) AND (($1::bigint << 55) | (l.line << 20) | 1048575) OFFSET 0) r "+
		"ORDER BY r.statement_id, r.ordinal", c.DumpID, pq.Array(lines))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []stmtRef
	for rows.Next() {
		var r stmtRef
		if err := rows.Scan(&r.Statement, &r.Ordinal, &r.Hash); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
