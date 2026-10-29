//ff:func feature=repair type=writer control=iteration dimension=1
//ff:what insertOrders — stores in a transaction the permutation of each occurrence in another order (reference_snak_order); a row already there takes the new one
package repair

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

// insertOrders stores each of orders in reference_snak_order in tx in one
// INSERT … SELECT FROM unnest: the occurrence (statement_id, ordinal) and its
// permutation, passed as an array literal per row (unnest flattens a
// two-dimensional array). An occurrence already stored (an overlapping or
// repeated range) keeps its row when the permutation is the same and takes
// the new one otherwise. An occurrence with no permutation (nil) is an error.
// Nothing for no orders.
func insertOrders(tx *sql.Tx, c Config, orders []OccOrder) error {
	if len(orders) == 0 {
		return nil
	}
	var sids, ords []int64
	var perms []string
	for _, o := range orders {
		if o.Perm == nil {
			return fmt.Errorf("reference %s at statement %d reference %d: no permutation", o.Hash, o.At.Statement, o.At.Ordinal)
		}
		sids, ords = append(sids, o.At.Statement), append(ords, int64(o.At.Ordinal))
		perms = append(perms, "{"+strings.Trim(strings.Join(strings.Fields(fmt.Sprint(o.Perm)), ","), "[]")+"}")
	}
	s := pq.QuoteIdentifier(c.Schema)
	_, err := tx.Exec("INSERT INTO "+s+".reference_snak_order AS o (statement_id, ordinal, perm) "+
		"SELECT st, od, p::smallint[] FROM unnest($1::bigint[], $2::smallint[], $3::text[]) AS t(st, od, p) "+
		"ON CONFLICT (statement_id, ordinal) DO UPDATE SET perm = EXCLUDED.perm WHERE o.perm IS DISTINCT FROM EXCLUDED.perm",
		pq.Array(sids), pq.Array(ords), pq.Array(perms))
	return err
}
