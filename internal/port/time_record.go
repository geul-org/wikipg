//ff:func feature=port type=encoder control=selection
//ff:what timeRecord — a GREGORIAN-DATE for a day-precision Gregorian Wikidata time, EDTF for a coarser one, GEDCOM-DATE for a Julian one; otherwise the reason it is skipped
package port

import (
	"encoding/json"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// timeRecord converts a timezone-0, unbounded Wikidata time datavalue. A
// Julian time is kept as written (julianRecord). Of a Gregorian one, a day
// (precision 11) is what GREGORIAN-DATE says exactly, so it is GREGORIAN-DATE
// (dayRecord), as the EDTF canonical form requires (value/EDTF). Month, year
// and decade precisions become an EDTF Value (edtfString) checked by go-edtf
// (edtfValid). Other calendars, timezones, uncertainty bounds, coarser
// precisions and strings go-edtf cannot check are skipped with their reason.
func timeRecord(value string) (core.Record, string) {
	var t wikidataTime
	if err := json.Unmarshal([]byte(value), &t); err != nil {
		return core.Record{}, "time: JSON 아님"
	}
	if t.Timezone != 0 || t.Before != 0 || t.After != 0 {
		return core.Record{}, "time: 시간대·오차 범위"
	}
	switch t.CalendarModel {
	case gregorian:
	case julian:
		return julianRecord(t)
	default:
		return core.Record{}, "time: " + calendarName(t.CalendarModel)
	}
	if t.Precision == 11 {
		return dayRecord(t.Time)
	}
	s, reason := edtfString(t)
	if reason != "" {
		return core.Record{}, reason
	}
	if err := edtfValid(s); err != nil {
		if strings.HasPrefix(s, "Y") {
			return core.Record{}, "time: EDTF 검증 불가(Y 접두 연도)"
		}
		return core.Record{}, "time: EDTF 검증 실패"
	}
	return core.EncodeBytesValue(core.CodecEDTFEven, []byte(s)), ""
}
