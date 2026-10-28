//ff:func feature=repair type=helper control=sequence
//ff:what TestVerdictString — each verdict has the name the done table stores
package repair

import "testing"

func TestVerdictString(t *testing.T) {
	if Same.String() != "same" || Order.String() != "order" || Content.String() != "content" || Missing.String() != "missing" {
		t.Fatal("names")
	}
}
