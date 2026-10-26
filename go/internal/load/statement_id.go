//ff:func feature=load type=helper control=sequence
//ff:what statementID — the statement key of the i-th statement on a dump line: dump << 55 | line << 20 | i
package load

import "fmt"

// statementID gives the i-th statement (0-based, properties in sorted order)
// of the entity on line line of dump dump its key, dump << 55 | line << 20 | i.
// It is decided by the dump alone, so parallel workers need no shared counter,
// a rerun gives the same keys and two dumps in one schema do not collide. The
// dump is 1..255, the line below 2^35 and an entity has at most 2^20
// statements.
func statementID(dump int16, line int64, i int) (int64, error) {
	if dump < 1 || dump > 255 || line < 0 || line >= 1<<35 || i < 0 || i >= 1<<20 {
		return 0, fmt.Errorf("load: statement %d of dump %d line %d out of key range", i, dump, line)
	}
	return int64(dump)<<55 | line<<20 | int64(i), nil
}
