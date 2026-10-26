//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestIsSegmentByte — pchar bytes pass; / ? # % [ ], space, non-ASCII and NUL fail
package port

import "testing"

func TestIsSegmentByte(t *testing.T) {
	for _, c := range []byte("aZ9-._~!$&'()*+,;=:@") {
		if !isSegmentByte(c) {
			t.Fatalf("%q: want allowed", c)
		}
	}
	for _, c := range []byte{'/', '?', '#', '%', '[', ']', ' ', 0xED, 0} {
		if isSegmentByte(c) {
			t.Fatalf("%q: want encoded", c)
		}
	}
}
