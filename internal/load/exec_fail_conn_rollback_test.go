//ff:func feature=load type=helper control=sequence
//ff:what execFailConn.Rollback — counts the rollback and succeeds
package load

// Rollback counts one rollback and succeeds.
func (c *execFailConn) Rollback() error {
	c.Rollbacks++
	return nil
}
