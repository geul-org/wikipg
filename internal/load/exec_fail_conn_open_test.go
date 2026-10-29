//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Open — the driver opens the same connection for any name
package load

import "database/sql/driver"

// Open returns c itself as the connection, whatever the name.
func (c *execFailConn) Open(string) (driver.Conn, error) { return c, nil }
