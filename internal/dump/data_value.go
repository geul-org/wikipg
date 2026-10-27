//ff:type feature=dump type=model
//ff:what DataValue — a snak's value: its datavalue type and the value JSON as written
package dump

import "encoding/json"

// DataValue is a snak's value: its type (wikibase-entityid, string, time,
// quantity, monolingualtext, globecoordinate …) and the value JSON exactly as
// the dump writes it (an entity {"entity-type":…,"id":"Q5"}, a quantity with
// its unit and bounds, a string …).
type DataValue struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}
