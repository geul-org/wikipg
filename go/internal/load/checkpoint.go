//ff:type feature=load type=model
//ff:what Checkpoint — tracks batches finishing out of order and the last line through which all are done
package load

// Checkpoint follows batches numbered from 0 that finish in any order and
// knows the last dump line through which every batch is done (Done). Line is
// that line; it starts at the line before the first batch.
type Checkpoint struct {
	next int64
	done map[int64]int64
	Line int64
}
