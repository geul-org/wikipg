//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what percentEncode — the bytes of s that keep does not allow as %XX (upper-case hex), UTF-8 bytes one by one
package port

import "strings"

// percentEncode writes each byte of s that keep rejects as %XX with upper-case
// hex digits (RFC 3986 2.1); a non-ASCII character becomes its UTF-8 bytes,
// each encoded (RFC 3987 3.1). Allowed bytes stay as they are.
func percentEncode(s string, keep func(byte) bool) string {
	const hex = "0123456789ABCDEF"
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if keep(c) {
			b.WriteByte(c)
			continue
		}
		b.WriteByte('%')
		b.WriteByte(hex[c>>4])
		b.WriteByte(hex[c&15])
	}
	return b.String()
}
