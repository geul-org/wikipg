//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Driver — the connector's driver is the connection itself
package load

import "database/sql/driver"

// Driver returns c itself as the driver.
func (c *execFailConn) Driver() driver.Driver { return c }
