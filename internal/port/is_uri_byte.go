//ff:func feature=port type=helper control=selection
//ff:what isURIByte — whether a byte may stand in a URI as is: unreserved, reserved or % (RFC 3986)
package port

// isURIByte reports whether c may appear in a URI unencoded (RFC 3986 2.2,
// 2.3): a letter, digit, - . _ ~, a reserved character : / ? # [ ] @ ! $ & '
// ( ) * + , ; = or % (an existing percent-encoding is kept).
func isURIByte(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
		return true
	default:
		return c != 0 && containsByte("-._~:/?#[]@!$&'()*+,;=%", c)
	}
}
