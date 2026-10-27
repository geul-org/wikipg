//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestWikibaseType — all 18 datatypes map to their ontology names; an unknown one is false
package port

import "testing"

func TestWikibaseType(t *testing.T) {
	for dt, want := range map[string]string{
		"wikibase-item": "WikibaseItem", "wikibase-property": "WikibaseProperty", "wikibase-lexeme": "WikibaseLexeme",
		"wikibase-form": "WikibaseForm", "wikibase-sense": "WikibaseSense", "entity-schema": "EntitySchema",
		"string": "String", "external-id": "ExternalId", "url": "Url", "commonsMedia": "CommonsMedia",
		"geo-shape": "GeoShape", "tabular-data": "TabularData", "monolingualtext": "Monolingualtext", "time": "Time",
		"globe-coordinate": "GlobeCoordinate", "quantity": "Quantity", "math": "Math", "musical-notation": "MusicalNotation",
	} {
		if got, ok := wikibaseType(dt); !ok || got != want {
			t.Fatalf("%s: %q %v", dt, got, ok)
		}
	}
	if _, ok := wikibaseType("nope"); ok {
		t.Fatal("unknown: want false")
	}
}
