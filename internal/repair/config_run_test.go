//ff:func feature=repair type=helper control=sequence
//ff:what TestConfigRun — a line range and an entity range get distinct checkpoint names
package repair

import "testing"

func TestConfigRun(t *testing.T) {
	if got := (Config{From: 2, To: 0}).Run(); got != "lines:2-0" {
		t.Fatal(got)
	}
	if got := (Config{Kind: "Q", Lo: 1, Hi: 10000}).Run(); got != "entities:Q1-Q10000" {
		t.Fatal(got)
	}
}
