//ff:func feature=port type=encoder control=sequence
//ff:what TestJulianRecord — a Julian day becomes GEDCOM-DATE; a century keeps gedcomString's reason
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestJulianRecord(t *testing.T) {
	must := func(b []byte, err error) []byte {
		if err != nil {
			t.Fatal(err)
		}
		return b
	}
	r, reason := julianRecord(wikidataTime{Time: "+1066-10-14T00:00:00Z", Precision: 11, CalendarModel: julian})
	if reason != "" || r.Codec != core.CodecGEDCOMDateEven || string(must(core.DecodeBytesValue(r, core.CodecGEDCOMDateEven))) != "JULIAN 14 OCT 1066" {
		t.Fatalf("%+v %q", r, reason)
	}
	if _, reason := julianRecord(wikidataTime{Time: "+1100-00-00T00:00:00Z", Precision: 7}); reason != "time: GEDCOM 대응 없는 정밀도 세기" {
		t.Fatalf("%q", reason)
	}
}
