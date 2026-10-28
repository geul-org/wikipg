//ff:func feature=repair type=helper control=sequence
//ff:what TestStatusLine — progress, verdicts and the time left at the pace so far; no batch yet gives no estimate
package repair

import (
	"strings"
	"testing"
	"time"
)

func TestStatusLine(t *testing.T) {
	s := statusLine(Stats{Spans: 1, Planned: 4, Last: 7, Entities: 10, Bytes: 2e6, Hashes: Tally{Order: 3}, Replaced: 3}, 10*time.Second)
	if !strings.Contains(s, "batch 1/4 line 7") || !strings.Contains(s, "order 3") || !strings.Contains(s, "left 30s") || !strings.Contains(s, "1 entities/s") {
		t.Fatal(s)
	}
	if s := statusLine(Stats{}, 0); !strings.Contains(s, "left 0s") {
		t.Fatal(s)
	}
}
