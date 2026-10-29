//ff:func feature=repair type=reader control=sequence
//ff:what checkKey — whether the run's checkpoint table exists, and an error when its key is not (dump_id, run) (a table made before 2026-10-29: sql/migrate_checkpoint_key.sql)
package repair

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// checkKey reports whether the checkpoint table of c's mode (Config.tables)
// exists in c.Schema, and fails when its primary key is not (dump_id, run).
// A table made before 2026-10-29 has the key run alone: saving under it would
// move another dump's line, so a run stops and names the migration instead.
func checkKey(db *sql.DB, c Config) (bool, error) {
	_, table := c.tables()
	name := pq.QuoteIdentifier(c.Schema) + "." + table
	var exists bool
	var key string
	err := db.QueryRow("SELECT to_regclass($1) IS NOT NULL, coalesce((SELECT string_agg(a.attname, ',' ORDER BY a.attname) "+
		"FROM pg_index i JOIN pg_attribute a ON a.attrelid = i.indrelid AND a.attnum = ANY (i.indkey) "+
		"WHERE i.indrelid = to_regclass($1) AND i.indisprimary), '')", name).Scan(&exists, &key)
	if err != nil || !exists {
		return false, err
	}
	if key != "dump_id,run" {
		return true, fmt.Errorf("repair: %s has the key (%s), not (dump_id, run): a checkpoint table from before 2026-10-29; "+
			"migrate it with PGOPTIONS='-c search_path=%s' psql -f sql/migrate_checkpoint_key.sql", name, key, c.Schema)
	}
	return true, nil
}
