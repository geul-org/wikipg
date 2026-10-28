//ff:func feature=repair type=helper control=sequence
//ff:what snakOfRow — a reference_snak row as the loader builds it, as a comparable Snak with the value in jsonb text form
package repair

// snakOfRow turns row (load.ReferenceSnaks: hash, dump_id, ordinal, property,
// value_ordinal, snaktype, datatype, value_type, value, value_entity,
// value_json; NULL as nil) into a Snak, the value replaced by its jsonb text
// form from norm (normalizeValues) so it compares with what the database
// prints.
func snakOfRow(row []any, norm map[string]string) Snak {
	s := Snak{Ordinal: row[2].(int), Property: row[3].(string), ValueOrdinal: row[4].(int), SnakType: row[5].(string),
		Datatype: nullText(row[6]), ValueType: nullText(row[7]), Value: nullText(row[8]), ValueEntity: nullText(row[9]), ValueJSON: nullText(row[10])}
	if s.Value.Valid {
		s.Value.String = norm[s.Value.String]
	}
	return s
}
