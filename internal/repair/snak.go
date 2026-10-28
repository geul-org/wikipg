//ff:type feature=repair type=model
//ff:what Snak — one reference_snak row as compared: ordinals, property, snak type and the value columns as text (value as jsonb text)
package repair

import "database/sql"

// Snak is one row of reference_snak without hash and dump_id, comparable with
// ==: overall and per-property ordinal, property, snak type, datatype, value
// type, value (jsonb text form, as PostgreSQL prints it), value entity and
// value_json. NULL is an invalid NullString.
type Snak struct {
	Ordinal      int
	ValueOrdinal int
	Property     string
	SnakType     string
	Datatype     sql.NullString
	ValueType    sql.NullString
	Value        sql.NullString
	ValueEntity  sql.NullString
	ValueJSON    sql.NullString
}
