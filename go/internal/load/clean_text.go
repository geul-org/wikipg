//ff:func feature=load type=helper control=sequence
//ff:what cleanText — a string without NUL bytes, which PostgreSQL text cannot hold; false when it had any
package load

import "strings"

// cleanText removes NUL bytes from s: PostgreSQL text and COPY reject them.
// It is false when s had one, so the caller records the change (the raw line
// keeps the original).
func cleanText(s string) (string, bool) {
	if !strings.ContainsRune(s, 0) {
		return s, true
	}
	return strings.ReplaceAll(s, "\x00", ""), false
}
