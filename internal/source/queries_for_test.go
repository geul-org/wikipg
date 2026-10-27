//ff:func feature=source type=loader control=iteration dimension=1
//ff:what TestQueriesFor — every query reads the given schema and dump and orders its rows by dump order
package source

import (
	"strings"
	"testing"
)

func TestQueriesFor(t *testing.T) {
	q := queriesFor(Store{Schema: "wd", Dump: 3})
	for _, s := range []string{q.Exists, q.Labels, q.Descriptions, q.Aliases, q.Claims, q.Qualifiers} {
		if !strings.Contains(s, "wd.") || !strings.Contains(s, "dump_id = 3") {
			t.Fatalf("schema or dump missing: %s", s)
		}
	}
	if !strings.Contains(q.Aliases, "ordinal") || !strings.Contains(q.Claims, "ORDER BY statement_id") || !strings.Contains(q.Qualifiers, "ORDER BY statement_id, ordinal") {
		t.Fatalf("order: %+v", q)
	}
}
