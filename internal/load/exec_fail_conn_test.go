//ff:type feature=load type=model
//ff:what execFailConn — a database/sql driver connection whose transactions begin but whose every statement fails, for error paths after Begin
package load

// execFailConn is a driver connector, driver, connection and transaction in
// one: Begin succeeds, every statement fails to prepare, and Rollbacks counts
// the rollbacks, so a test sees that a failed statement rolls back.
type execFailConn struct {
	Rollbacks int
}
