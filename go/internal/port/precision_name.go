//ff:func feature=port type=helper control=selection
//ff:what precisionName — a readable name of a Wikidata time precision for the skip report
package port

import "strconv"

// precisionName names a Wikidata time precision for the skip report.
func precisionName(p int) string {
	switch p {
	case 10:
		return "월"
	case 9:
		return "연"
	case 8:
		return "10년"
	case 7:
		return "세기"
	case 6:
		return "천년"
	default:
		return strconv.Itoa(p)
	}
}
