//ff:func feature=port type=encoder control=selection
//ff:what edtfString — the EDTF string of a Wikidata time at day, month, year or decade precision; otherwise the reason it is skipped
package port

import "fmt"

// edtfString writes the date of a Gregorian, timezone-0, unbounded Wikidata
// time as EDTF (ISO 8601-2): day 1952-03-11, month 1952-03, year 1952, decade
// 195X. Coarser precisions are not mapped yet (how Wikidata centuries map to
// EDTF unspecified digits is open). The caller checks
// calendar, timezone and bounds.
func edtfString(t wikidataTime) (string, string) {
	y, m, d, ok := parseCivilDate(t.Time)
	if !ok {
		return "", "time: 날짜 형식"
	}
	year, ok := edtfYear(y)
	if !ok {
		return "", "time: 0년"
	}
	switch t.Precision {
	case 11:
		return fmt.Sprintf("%s-%02d-%02d", year, m, d), ""
	case 10:
		return fmt.Sprintf("%s-%02d", year, m), ""
	case 9:
		return year, ""
	case 8:
		if len(year) != 4 {
			return "", "time: 10년(기원전·큰 연도)"
		}
		return year[:3] + "X", ""
	default:
		return "", "time: EDTF 대응 미결정 정밀도 " + precisionName(t.Precision)
	}
}
