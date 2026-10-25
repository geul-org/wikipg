//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestDayRecord — CE and BCE days convert (BCE shifted to astronomical years); malformed, year 0, impossible days and out-of-range years are skipped
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDayRecord(t *testing.T) {
	r, reason := dayRecord("+1970-01-02T00:00:00Z")
	if reason != "" || !slices.Equal(r.Payload, []core.Word{0, 1}) {
		t.Fatalf("CE: %v %q", r.Payload, reason)
	}
	r, reason = dayRecord("-0001-01-01T00:00:00Z")
	if d := int32(uint32(r.Payload[0])<<16 | uint32(r.Payload[1])); reason != "" || int64(d) != daysFromCivil(0, 1, 1) {
		t.Fatalf("1 BCE: %d %q", d, reason)
	}
	for want, ts := range map[string]string{
		"time: 날짜 형식": "x",
		"time: 범위":    "+9999999-01-01T00:00:00Z",
	} {
		if _, got := dayRecord(ts); got != want {
			t.Fatalf("%s: %q want %q", ts, got, want)
		}
	}
	for _, ts := range []string{"+0000-01-01T00:00:00Z", "+2023-02-29T00:00:00Z"} {
		if _, got := dayRecord(ts); got != "time: 날짜 형식" {
			t.Fatalf("%s: %q", ts, got)
		}
	}
}
