//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what windowSpans — line windows of batch lines covering from..to, numbered in order
package repair

// windowSpans cuts lines from..to into windows of batch lines, numbered
// from 0; each window's last line is its checkpoint.
func windowSpans(from, to int64, batch int) []Span {
	var out []Span
	for a := from; a <= to; a += int64(batch) {
		b := min(a+int64(batch)-1, to)
		out = append(out, Span{Seq: int64(len(out)), From: a, To: b, Last: b})
	}
	return out
}
