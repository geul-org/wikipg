//ff:func feature=source type=loader control=sequence
//ff:what TestDefaultQueries — every query is set and orders its rows (qualifiers by physical row order)
package source

import (
	"strings"
	"testing"
)

func TestDefaultQueries(t *testing.T) {
	q := defaultQueries()
	if q.Exists == "" || !strings.Contains(q.Labels, "ORDER BY") || !strings.Contains(q.Descriptions, "entity_descriptions") || !strings.Contains(q.Aliases, "alias_order") || !strings.Contains(q.Claims, "ORDER BY id") || !strings.Contains(q.Qualifiers, "ctid") {
		t.Fatalf("got %+v", q)
	}
}
