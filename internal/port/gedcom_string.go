//ff:func feature=port type=encoder control=selection
//ff:what gedcomString — the GEDCOM 7 date value of a Julian Wikidata time (day, month or year precision), or the reason it is skipped
package port

import "fmt"

// gedcomString writes the date of a Julian, timezone-0, unbounded Wikidata
// time as a GEDCOM 7 date value with the JULIAN calendar keyword: day
// "JULIAN 14 OCT 1066", month "JULIAN OCT 1066", year "JULIAN 1066". Wikidata
// years are historical (-0044 is 44 BCE) like GEDCOM, so a BCE year keeps its
// number and takes the BCE epoch. The day must exist in the Julian calendar
// (validJulianDay). GEDCOM has no decade or century form, so coarser
// precisions are skipped.
func gedcomString(t wikidataTime) (string, string) {
	y, m, d, ok := parseCivilDate(t.Time)
	if !ok || y == 0 {
		return "", "time: 날짜 형식"
	}
	year := fmt.Sprint(y)
	if y < 0 {
		year = fmt.Sprint(-y) + " BCE"
	}
	months := [13]string{"", "JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	switch {
	case t.Precision == 11 && validJulianDay(y, m, d):
		return fmt.Sprintf("JULIAN %d %s %s", d, months[m], year), ""
	case t.Precision == 10 && m >= 1 && m <= 12:
		return fmt.Sprintf("JULIAN %s %s", months[m], year), ""
	case t.Precision == 9:
		return "JULIAN " + year, ""
	case t.Precision == 11 || t.Precision == 10:
		return "", "time: 날짜 형식"
	default:
		return "", "time: GEDCOM 대응 없는 정밀도 " + precisionName(t.Precision)
	}
}
