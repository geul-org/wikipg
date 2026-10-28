//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what listSpans — a sorted list of lines in batches of batch lines, numbered in order
package repair

// listSpans cuts the sorted lines into batches of batch lines, numbered from
// 0; each batch's last line is its checkpoint.
func listSpans(lines []int64, batch int) []Span {
	var out []Span
	for a := 0; a < len(lines); a += batch {
		part := lines[a:min(a+batch, len(lines))]
		out = append(out, Span{Seq: int64(len(out)), Lines: part, Last: part[len(part)-1]})
	}
	return out
}
