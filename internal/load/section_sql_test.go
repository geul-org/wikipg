//ff:func feature=load type=helper control=iteration dimension=1
//ff:what sectionSQL — the commented SQL block of one section of sql/wikidata.sql section 4 (BEGIN … COMMIT), uncommented and set to a schema
package load

import (
	"os"
	"strings"
	"testing"
)

// sectionSQL returns the SQL of section (e.g. "-- 4.2 ") of sql/wikidata.sql:
// the commented lines from its first "-- BEGIN;" to the next "-- COMMIT;"
// without the "-- " prefix, with the schema wikidata replaced by schema. The
// test runs the very text the operator runs.
func sectionSQL(t *testing.T, section, schema string) string {
	t.Helper()
	b, err := os.ReadFile("../../sql/wikidata.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, rest, ok := strings.Cut(string(b), "\n"+section)
	if !ok {
		t.Fatalf("no section %q", section)
	}
	var out []string
	in := false
	for _, line := range strings.Split(rest, "\n") {
		in = in || line == "-- BEGIN;"
		if in {
			out = append(out, strings.TrimPrefix(line, "-- "))
		}
		if in && line == "-- COMMIT;" {
			break
		}
	}
	return strings.Replace(strings.Join(out, "\n"), "SET search_path = wikidata;", "SET search_path = "+schema+";", 1)
}
