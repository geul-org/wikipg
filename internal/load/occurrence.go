//ff:type feature=load type=model
//ff:what Occurrence — where a reference occurs in the dump: the statement key and its place in that statement's references
package load

// Occurrence is one place a reference (hash) occurs: the key of the statement
// that cites it (statementID: dump << 55 | line << 20 | i, so keys follow dump
// order) and its position in that statement's references (0-based). The
// smallest Occurrence of a hash (Before) is its first occurrence in dump
// order, whose snak list the hash keeps (sql/wikidata.sql 4.2).
type Occurrence struct {
	Statement int64
	Ordinal   int
}
