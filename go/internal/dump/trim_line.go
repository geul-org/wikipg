//ff:func feature=dump type=decoder control=sequence
//ff:what TrimLine — a dump line without its trailing comma and newline; false for the array brackets and blank lines
package dump

import "bytes"

// TrimLine returns the entity JSON of one dump line: the line without the
// newline, carriage return and the comma that separates array elements. The
// dump is one JSON array with one entity per line, so the "[" and "]" lines
// and blank lines carry no entity and are false.
func TrimLine(line []byte) ([]byte, bool) {
	b := bytes.TrimRight(line, "\r\n")
	b = bytes.TrimSuffix(b, []byte(","))
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("[")) || bytes.Equal(b, []byte("]")) {
		return nil, false
	}
	return b, true
}
