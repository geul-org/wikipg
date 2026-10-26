//ff:type feature=dump type=model
//ff:what Snak — one property-value pair: snak type, property, hash, datatype, datavalue
package dump

// Snak is a main snak, qualifier or reference snak: its type (value, novalue,
// somevalue), property, hash, the property's datatype and, for a value snak,
// the datavalue.
type Snak struct {
	SnakType  string     `json:"snaktype"`
	Property  string     `json:"property"`
	Hash      string     `json:"hash"`
	Datatype  string     `json:"datatype"`
	DataValue *DataValue `json:"datavalue"`
}
