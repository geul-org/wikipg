//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestTimeRecord — days become GREGORIAN-DATE, coarser precisions EDTF; Julian times become GEDCOM-DATE; other calendars, bounds, coarse precisions and unverifiable EDTF are skipped with reasons
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestTimeRecord(t *testing.T) {
	tv := func(ts, prec, cal string) string {
		return `{"time":"` + ts + `","timezone":0,"before":0,"after":0,"precision":` + prec + `,"calendarmodel":"` + cal + `"}`
	}
	enc := func(v string) []core.Word {
		r, reason := timeRecord(v)
		if reason != "" {
			t.Fatalf("%s: %q", v, reason)
		}
		ws, _ := core.EncodeRecord(r)
		return ws
	}
	if ws := enc(tv("+1952-03-11T00:00:00Z", "11", gregorian)); !slices.Equal(ws, []core.Word{0x0920, 0xFFFF, 0xE697}) {
		t.Fatalf("day: %04X", ws)
	}
	if ws := enc(tv("+1974-00-00T00:00:00Z", "9", gregorian)); !slices.Equal(ws, []core.Word{0x0924, 0x3139, 0x3734}) {
		t.Fatalf("year: %04X", ws)
	}
	if ws := enc(tv("+1300-00-00T00:00:00Z", "9", julian)); !slices.Equal(ws, []core.Word{0x0B27, 0x4A55, 0x4C49, 0x414E, 0x2031, 0x3330, 0x3000}) {
		t.Fatalf("julian: %04X", ws)
	}
	for want, v := range map[string]string{
		"time: JSON 아님":             "x",
		"time: GEDCOM 대응 없는 정밀도 세기": tv("+1300-00-00T00:00:00Z", "7", julian),
		"time: 달력 x":                tv("+1300-01-01T00:00:00Z", "11", "x"),
		"time: 시간대·오차 범위":           `{"time":"+1952-03-11T00:00:00Z","timezone":60,"precision":11,"calendarmodel":"` + gregorian + `"}`,
		"time: EDTF 대응 미결정 정밀도 세기":  tv("+1900-00-00T00:00:00Z", "7", gregorian),
		"time: EDTF 검증 불가(Y 접두 연도)": tv("+12000-00-00T00:00:00Z", "9", gregorian),
		"time: EDTF 검증 실패":          tv("+1974-13-00T00:00:00Z", "10", gregorian),
		"time: 날짜 형식":               tv("+2023-02-29T00:00:00Z", "11", gregorian),
	} {
		if _, got := timeRecord(v); got != want {
			t.Fatalf("%s: %q want %q", v, got, want)
		}
	}
}
