//ff:func feature=port type=helper control=sequence
//ff:what parseCivilDate — year, month and day of a Wikidata timestamp such as +1952-03-11T00:00:00Z; false when malformed
package port

import (
	"strconv"
	"strings"
)

// parseCivilDate reads the date part of a Wikidata timestamp:
// sign, year (any digits), month and day, then "T00:00:00Z". Month and day
// may be 00 for coarser precisions; this returns them as they are.
func parseCivilDate(ts string) (year int64, month, day int, ok bool) {
	if len(ts) < 1 || (ts[0] != '+' && ts[0] != '-') || !strings.HasSuffix(ts, "T00:00:00Z") {
		return 0, 0, 0, false
	}
	parts := strings.Split(strings.TrimSuffix(ts[1:], "T00:00:00Z"), "-")
	if len(parts) != 3 || len(parts[1]) != 2 || len(parts[2]) != 2 {
		return 0, 0, 0, false
	}
	y, err1 := strconv.ParseInt(parts[0], 10, 64)
	m, err2 := strconv.Atoi(parts[1])
	d, err3 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, false
	}
	if ts[0] == '-' {
		y = -y
	}
	return y, m, d, true
}
