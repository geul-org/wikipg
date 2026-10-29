//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Prepare — every statement fails to prepare
package load

import (
	"database/sql/driver"
	"errors"
)

// Prepare fails for every query, so every Exec on the connection fails.
func (c *execFailConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("statement refused")
}
