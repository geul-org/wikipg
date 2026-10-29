//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Close — closing the connection always succeeds
package load

// Close does nothing and succeeds.
func (c *execFailConn) Close() error { return nil }
