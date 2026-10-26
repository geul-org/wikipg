//ff:func feature=dump type=decoder control=iteration dimension=1
//ff:what TestValueEntity — an entity datavalue gives its ID; nil, other types, bad JSON and no id give empty
package dump

import (
	"encoding/json"
	"testing"
)

func TestValueEntity(t *testing.T) {
	if got := ValueEntity(&DataValue{Type: "wikibase-entityid", Value: json.RawMessage(`{"entity-type":"item","id":"Q5"}`)}); got != "Q5" {
		t.Fatalf("got %q", got)
	}
	for _, dv := range []*DataValue{nil, {Type: "string", Value: json.RawMessage(`"Q5"`)}, {Type: "wikibase-entityid", Value: json.RawMessage(`{`)}, {Type: "wikibase-entityid", Value: json.RawMessage(`{}`)}} {
		if got := ValueEntity(dv); got != "" {
			t.Fatalf("%+v: %q", dv, got)
		}
	}
}
