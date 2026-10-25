//ff:func feature=port type=helper control=sequence
//ff:what entityNumber — the prefix letter and number of a Wikidata ID (Q42 → Q, 42); false when it is not one letter and a positive number
package port

import "strconv"

// entityNumber splits a Wikidata ID such as Q42 or P31 into its prefix letter
// and positive number.
func entityNumber(id string) (byte, uint64, bool) {
	if len(id) < 2 {
		return 0, 0, false
	}
	n, err := strconv.ParseUint(id[1:], 10, 64)
	if err != nil || n == 0 || id[1] == '0' || id[1] == '+' {
		return 0, 0, false
	}
	return id[0], n, true
}
