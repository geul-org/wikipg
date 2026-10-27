//ff:func feature=load type=helper control=sequence
//ff:what TestCleanText — text without NUL stays and is true; NUL bytes are removed and false
package load

import "testing"

func TestCleanText(t *testing.T) {
	if s, ok := cleanText("가a"); !ok || s != "가a" {
		t.Fatalf("%q %v", s, ok)
	}
	if s, ok := cleanText("a\x00b\x00"); ok || s != "ab" {
		t.Fatalf("%q %v", s, ok)
	}
}
