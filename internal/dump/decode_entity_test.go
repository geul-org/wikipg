//ff:func feature=dump type=decoder control=iteration dimension=1
//ff:what TestDecodeEntity — the sample line decodes with revision, terms, claims, qualifier order, references and novalue; bad JSON and a missing ID fail
package dump

import "testing"

func TestDecodeEntity(t *testing.T) {
	raw, _ := TrimLine(sampleLine())
	e, err := DecodeEntity(raw)
	if err != nil || e.ID != "Q42" || e.Type != "item" || e.LastRevID != 2401234567 || e.Modified != "2026-10-18T09:00:00Z" {
		t.Fatalf("%+v %v", e, err)
	}
	if e.Labels["ko"].Value != "더글러스 애덤스" || len(e.Aliases["en"]) != 2 || e.Descriptions["en"].Value != "English writer" || e.Sitelinks["enwiki"].Badges[0] != "Q17437796" {
		t.Fatalf("terms %+v", e)
	}
	s := e.Claims["P69"][0]
	if s.ID != "Q42$A" || s.QualifiersOrder[0] != "P580" || len(s.Qualifiers["P582"]) != 1 || s.References[0].Hash != "r1" || s.References[0].SnaksOrder[0] != "P248" {
		t.Fatalf("claim %+v", s)
	}
	if n := e.Claims["P40"][0].Mainsnak; n.SnakType != "novalue" || n.DataValue != nil {
		t.Fatalf("novalue %+v", n)
	}
	if string(e.Claims["P2048"][0].Mainsnak.DataValue.Value) != `{"amount":"+1.96","unit":"http://www.wikidata.org/entity/Q11573","upperBound":"+1.97","lowerBound":"+1.95"}` {
		t.Fatal("quantity value not kept as written")
	}
	for _, bad := range []string{"[", `{"type":"item"}`} {
		if _, err := DecodeEntity([]byte(bad)); err == nil {
			t.Fatalf("%q: want error", bad)
		}
	}
}
