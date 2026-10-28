//ff:type feature=repair type=model
//ff:what snakColumns — reference_snak rows laid out by column, as arrays for one INSERT … SELECT FROM unnest
package repair

import "database/sql"

// snakColumns holds rows of reference_snak column by column (hash, dump_id,
// ordinal, property, value_ordinal, snaktype, datatype, value_type, value,
// value_entity, value_json), NULL as an invalid NullString, so many rows go
// in one statement.
type snakColumns struct {
	Hash, Property, SnakType                           []string
	DumpID, Ordinal, ValueOrdinal                      []int64
	Datatype, ValueType, Value, ValueEntity, ValueJSON []sql.NullString
}
