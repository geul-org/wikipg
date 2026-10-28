//ff:type feature=repair type=model
//ff:what stmtRef — one statement_reference row: statement key, position, hash
package repair

// stmtRef is one row of statement_reference: the statement key, the
// position in its references and the reference hash.
type stmtRef struct {
	Statement int64
	Ordinal   int
	Hash      string
}
