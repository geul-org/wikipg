//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Commit — committing always succeeds
package load

// Commit does nothing and succeeds.
func (c *execFailConn) Commit() error { return nil }
