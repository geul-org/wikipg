//ff:func feature=port type=helper control=sequence
//ff:what containsByte — whether a byte is one of the bytes of a set string
package port

import "strings"

// containsByte reports whether c is one of the bytes of set.
func containsByte(set string, c byte) bool {
	return strings.IndexByte(set, c) >= 0
}
