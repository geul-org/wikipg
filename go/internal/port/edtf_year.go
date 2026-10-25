//ff:func feature=port type=helper control=selection
//ff:what edtfYear — the EDTF year of a Wikidata year: BCE shifted to astronomical numbering, four digits or a Y prefix
package port

import "fmt"

// edtfYear writes year y of a Wikidata timestamp as an EDTF year. Wikidata
// numbers BCE years historically (-0044 is 44 BCE, there is no year 0) while
// EDTF follows ISO 8601 astronomical numbering (0000 is 1 BCE, -0043 is 44
// BCE), so a negative year moves up by one. Years outside ±9999 take the EDTF
// Y prefix. Year 0 does not exist in Wikidata and is false.
func edtfYear(y int64) (string, bool) {
	a := y
	if y < 0 {
		a = y + 1
	}
	switch {
	case y == 0:
		return "", false
	case a > 9999 || a < -9999:
		return fmt.Sprintf("Y%d", a), true
	case a < 0:
		return fmt.Sprintf("-%04d", -a), true
	default:
		return fmt.Sprintf("%04d", a), true
	}
}
