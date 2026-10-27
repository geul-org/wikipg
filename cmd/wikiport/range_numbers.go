//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what rangeNumbers — the numbers from through to, inclusive
package main

// rangeNumbers returns from, from+1, …, to.
func rangeNumbers(from, to uint64) []uint64 {
	var out []uint64
	for n := from; n <= to; n++ {
		out = append(out, n)
	}
	return out
}
