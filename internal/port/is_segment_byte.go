//ff:func feature=port type=helper control=selection
//ff:what isSegmentByte — whether a byte may stand in one URI path segment as is (RFC 3986 pchar without %)
package port

// isSegmentByte reports whether c may appear unencoded in a single path
// segment (RFC 3986 3.3 pchar): a letter, digit, - . _ ~, ! $ & ' ( ) * + , ;
// = : or @. / ? # % and everything else are encoded, so a title such as
// "What?.jpg" stays one segment.
func isSegmentByte(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
		return true
	default:
		return c != 0 && containsByte("-._~!$&'()*+,;=:@", c)
	}
}
