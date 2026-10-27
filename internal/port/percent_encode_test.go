//ff:func feature=port type=encoder control=sequence
//ff:what TestPercentEncode — allowed bytes stay, others and each UTF-8 byte become upper-case %XX
package port

import "testing"

func TestPercentEncode(t *testing.T) {
	if got := percentEncode("a b/한", isURIByte); got != "a%20b/%ED%95%9C" {
		t.Fatalf("got %q", got)
	}
	if got := percentEncode("a/b?c", isSegmentByte); got != "a%2Fb%3Fc" {
		t.Fatalf("segment: %q", got)
	}
}
