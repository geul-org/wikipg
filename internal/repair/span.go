//ff:type feature=repair type=model
//ff:what Span — one batch of dump lines: its number in order, a line window or a list of lines, and the last line (the checkpoint after it)
package repair

// Span is one batch: Seq numbers batches from 0 in line order; Lines lists
// the lines (entity ranges) or, when nil, From..To is the window (line
// ranges); Last is the last line it covers, where the checkpoint moves once
// it is applied.
type Span struct {
	Seq   int64
	From  int64
	To    int64
	Lines []int64
	Last  int64
}
