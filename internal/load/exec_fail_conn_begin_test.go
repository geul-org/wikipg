//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Begin — a transaction always begins, as the connection itself
package load

import "database/sql/driver"

// Begin succeeds and returns c itself as the transaction.
func (c *execFailConn) Begin() (driver.Tx, error) { return c, nil }
