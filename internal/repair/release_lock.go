//ff:func feature=repair type=writer control=sequence
//ff:what releaseLock — releases the repair's advisory lock and its connection
package repair

import (
	"context"
	"database/sql"
)

// releaseLock releases the advisory lock of c.Schema taken by takeLock and
// returns its connection to the pool.
func releaseLock(conn *sql.Conn, c Config) error {
	_, err := conn.ExecContext(context.Background(), "SELECT pg_advisory_unlock(hashtext($1))", lockKey+" "+c.Schema)
	if cerr := conn.Close(); err == nil {
		err = cerr
	}
	return err
}
