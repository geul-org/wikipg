//ff:func feature=dump type=decoder control=sequence
//ff:what ValueEntity — the entity ID a datavalue points at; empty for other values
package dump

import "encoding/json"

// ValueEntity returns the "id" of a wikibase-entityid datavalue (Q5, P31,
// L7-S1 …) and "" for any other value, a missing value or one without an id.
func ValueEntity(dv *DataValue) string {
	if dv == nil || dv.Type != "wikibase-entityid" {
		return ""
	}
	var v struct {
		ID string `json:"id"`
	}
	if json.Unmarshal(dv.Value, &v) != nil {
		return ""
	}
	return v.ID
}
