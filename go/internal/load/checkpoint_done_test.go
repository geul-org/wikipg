//ff:func feature=load type=model control=sequence
//ff:what TestCheckpointDone — out-of-order batches advance the line only when the gap closes
package load

import "testing"

func TestCheckpointDone(t *testing.T) {
	c := NewCheckpoint(1)
	if c.Done(1, 20) || c.Line != 1 {
		t.Fatal("batch 1 before 0 must not move")
	}
	if !c.Done(0, 10) || c.Line != 20 {
		t.Fatalf("after 0: %d", c.Line)
	}
	if !c.Done(2, 30) || c.Line != 30 || len(c.done) != 0 {
		t.Fatalf("after 2: %d", c.Line)
	}
}
