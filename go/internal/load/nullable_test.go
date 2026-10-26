//ff:func feature=load type=helper control=sequence
//ff:what TestNullable — empty is NULL, anything else itself
package load

import "testing"

func TestNullable(t *testing.T) {
	if nullable("") != nil || nullable("x") != "x" {
		t.Fatal("wrong")
	}
}
