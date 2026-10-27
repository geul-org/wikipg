//ff:func feature=load type=model control=sequence
//ff:what TestNewCheckpoint — starts at the given line expecting batch 0
package load

import "testing"

func TestNewCheckpoint(t *testing.T) {
	if c := NewCheckpoint(42); c.Line != 42 || c.next != 0 || c.done == nil {
		t.Fatalf("got %+v", c)
	}
}
