//ff:func feature=port type=encoder control=sequence
//ff:what dayRecord — the GREGORIAN-DATE of a day-precision Wikidata timestamp, BCE included; otherwise the reason it is skipped
package port

import (
	"math"

	"github.com/geul-org/gogeul/pkg/core"
)

// dayRecord converts the date of a day-precision Wikidata timestamp to
// GREGORIAN-DATE (proleptic Gregorian days from 1970-01-01). A BCE year is
// historical in Wikidata (-0044 is 44 BCE, no year 0) and moves up by one to
// the astronomical year (-43) first. The day must exist and fit int32 days.
func dayRecord(ts string) (core.Record, string) {
	y, m, d, ok := parseCivilDate(ts)
	if !ok || y == 0 {
		return core.Record{}, "time: 날짜 형식"
	}
	if y < 0 {
		y++
	}
	if !validDay(y, m, d) {
		return core.Record{}, "time: 날짜 형식"
	}
	days := daysFromCivil(y, m, d)
	if days < math.MinInt32 || days > math.MaxInt32 {
		return core.Record{}, "time: 범위"
	}
	u := uint32(int32(days))
	return core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{core.Word(u >> 16), core.Word(u)}}, ""
}
