//ff:func feature=port type=helper control=sequence
//ff:what TestDocNumber — Q42 and P31 put their letter in the top byte; a number of 2^48 fails
package port

import "testing"

func TestDocNumber(t *testing.T) {
	if n, err := DocNumber('Q', 42); err != nil || n != 0x5100_0000_0000_002A {
		t.Fatalf("Q42: %x %v", n, err)
	}
	if n, err := DocNumber('P', 31); err != nil || n != 0x5000_0000_0000_001F {
		t.Fatalf("P31: %x %v", n, err)
	}
	if _, err := DocNumber('Q', 1<<56); err == nil {
		t.Fatal("2^56: want error")
	}
}
