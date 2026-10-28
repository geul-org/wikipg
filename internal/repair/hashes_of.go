//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what hashesOf — the hashes of some checks, in order
package repair

// hashesOf returns the hash of each of checks in order.
func hashesOf(checks []HashCheck) []string {
	out := make([]string, 0, len(checks))
	for _, h := range checks {
		out = append(out, h.Hash)
	}
	return out
}
