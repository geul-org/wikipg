//ff:func feature=port type=helper control=sequence
//ff:what TestContainsByte — a byte in the set is found, one outside and any in an empty set are not
package port

import "testing"

func TestContainsByte(t *testing.T) {
	if !containsByte("ab", 'b') || containsByte("ab", 'c') || containsByte("", 'a') {
		t.Fatal("containsByte")
	}
}
