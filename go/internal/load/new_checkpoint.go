//ff:func feature=load type=model control=sequence
//ff:what NewCheckpoint — a Checkpoint whose line before the first batch is given
package load

// NewCheckpoint returns a Checkpoint for batches numbered from 0, the first
// starting after line.
func NewCheckpoint(line int64) *Checkpoint {
	return &Checkpoint{done: map[int64]int64{}, Line: line}
}
