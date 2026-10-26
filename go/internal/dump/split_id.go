//ff:func feature=dump type=decoder control=sequence
//ff:what SplitID — an entity ID's letter and number (Q42 → Q, 42); false for anything else
package dump

import "strconv"

// SplitID splits an entity ID such as Q42, P31 or L7 into its letter and
// positive number. A leading zero, a sign or a non-number is false.
func SplitID(id string) (byte, int64, bool) {
	if len(id) < 2 || id[1] < '1' || id[1] > '9' {
		return 0, 0, false
	}
	n, err := strconv.ParseInt(id[1:], 10, 64)
	if err != nil {
		return 0, 0, false
	}
	return id[0], n, true
}
