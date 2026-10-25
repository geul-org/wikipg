//ff:func feature=source type=loader control=sequence
//ff:what TestDefaultQueries — every query is set and orders its rows
package source

import (
	"strings"
	"testing"
)

func TestDefaultQueries(t *testing.T) {
	q := defaultQueries()
	if q.Exists == "" || !strings.Contains(q.Labels, "ORDER BY") || !strings.Contains(q.Aliases, "alias_order") || !strings.Contains(q.Claims, "ORDER BY id") {
		t.Fatalf("got %+v", q)
	}
}
