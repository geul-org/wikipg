//ff:func feature=port type=encoder control=sequence
//ff:what julianRecord — a GEDCOM-DATE Value for a Julian Wikidata time, or the reason it is not converted
package port

import "github.com/geul-org/geul/go/pkg/core"

// julianRecord keeps a Julian Wikidata time as written: a GEDCOM-DATE Value
// (gedcomString) checked by gedcom-go (gedcomValid). The Gregorian value of
// the same day or period is not written yet: how a TRIPLE carries both
// values is open (docs/value/time/GEDCOM-DATE.md).
func julianRecord(t wikidataTime) (core.Record, string) {
	s, reason := gedcomString(t)
	if reason != "" {
		return core.Record{}, reason
	}
	if err := gedcomValid(s); err != nil {
		return core.Record{}, "time: GEDCOM 검증 실패"
	}
	return core.EncodeBytesValue(core.CodecGEDCOMDateEven, []byte(s)), ""
}
