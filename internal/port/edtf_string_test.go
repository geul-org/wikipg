//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestEdtfString — day, month, year and decade map to EDTF; malformed, year 0, BCE decades and coarser precisions are skipped
package port

import "testing"

func TestEdtfString(t *testing.T) {
	for want, tv := range map[string]wikidataTime{
		"1952-03-11": {Time: "+1952-03-11T00:00:00Z", Precision: 11},
		"1952-03":    {Time: "+1952-03-00T00:00:00Z", Precision: 10},
		"-0043":      {Time: "-0044-00-00T00:00:00Z", Precision: 9},
		"197X":       {Time: "+1970-00-00T00:00:00Z", Precision: 8},
	} {
		if got, reason := edtfString(tv); got != want || reason != "" {
			t.Fatalf("%+v: %q %q want %q", tv, got, reason, want)
		}
	}
	for want, tv := range map[string]wikidataTime{
		"time: 날짜 형식":              {Time: "x", Precision: 11},
		"time: 0년":                 {Time: "+0000-00-00T00:00:00Z", Precision: 9},
		"time: 10년(기원전·큰 연도)":      {Time: "-0500-00-00T00:00:00Z", Precision: 8},
		"time: EDTF 대응 미결정 정밀도 세기": {Time: "+1900-00-00T00:00:00Z", Precision: 7},
	} {
		if _, got := edtfString(tv); got != want {
			t.Fatalf("%+v: %q want %q", tv, got, want)
		}
	}
}
