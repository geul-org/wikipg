//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestGedcomString — day, month, year and BCE Julian times; malformed, year 0, impossible days and months, coarse precisions are skipped
package port

import "testing"

func TestGedcomString(t *testing.T) {
	tm := func(ts string, p int) wikidataTime {
		return wikidataTime{Time: ts, Precision: p, CalendarModel: julian}
	}
	for want, in := range map[string]wikidataTime{
		"JULIAN 14 OCT 1066":   tm("+1066-10-14T00:00:00Z", 11),
		"JULIAN OCT 1066":      tm("+1066-10-00T00:00:00Z", 10),
		"JULIAN 1066":          tm("+1066-00-00T00:00:00Z", 9),
		"JULIAN 15 MAR 44 BCE": tm("-0044-03-15T00:00:00Z", 11),
		"JULIAN 29 FEB 1100":   tm("+1100-02-29T00:00:00Z", 11),
	} {
		if got, reason := gedcomString(in); got != want || reason != "" {
			t.Fatalf("%v: %q %q", in, got, reason)
		}
	}
	skips := []struct {
		in   wikidataTime
		want string
	}{
		{tm("x", 9), "time: 날짜 형식"},
		{tm("+0000-00-00T00:00:00Z", 9), "time: 날짜 형식"},
		{tm("+1100-02-30T00:00:00Z", 11), "time: 날짜 형식"},
		{tm("+1100-13-00T00:00:00Z", 10), "time: 날짜 형식"},
		{tm("+1060-00-00T00:00:00Z", 8), "time: GEDCOM 대응 없는 정밀도 10년"},
	}
	for _, c := range skips {
		if _, got := gedcomString(c.in); got != c.want {
			t.Fatalf("%v: %q want %q", c.in, got, c.want)
		}
	}
}
