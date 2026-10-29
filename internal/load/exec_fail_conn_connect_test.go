//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Connect — the connector hands out the same connection
package load

import (
	"context"
	"database/sql/driver"
)

// Connect returns c itself as the connection.
func (c *execFailConn) Connect(context.Context) (driver.Conn, error) { return c, nil }
