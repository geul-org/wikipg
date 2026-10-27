//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestIsURIByte — letters, digits, unreserved, reserved and % pass; space, quote, backslash, non-ASCII and NUL fail
package port

import "testing"

func TestIsURIByte(t *testing.T) {
	for _, c := range []byte("aZ9-._~:/?#[]@!$&'()*+,;=%") {
		if !isURIByte(c) {
			t.Fatalf("%q: want allowed", c)
		}
	}
	for _, c := range []byte{' ', '"', '<', '\\', 0xED, 0} {
		if isURIByte(c) {
			t.Fatalf("%q: want encoded", c)
		}
	}
}
