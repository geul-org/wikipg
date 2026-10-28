//ff:func feature=repair type=helper control=iteration dimension=2
//ff:what columnsOf — the rows of some checks laid out by column (snakColumns)
package repair

// columnsOf lays out the Rows of checks (load.ReferenceSnaks: hash, dump_id
// as int16, ordinal and value_ordinal as int, text or nil) column by column.
func columnsOf(checks []HashCheck) snakColumns {
	var c snakColumns
	for _, h := range checks {
		for _, r := range h.Rows {
			c.Hash, c.DumpID, c.Ordinal = append(c.Hash, r[0].(string)), append(c.DumpID, int64(r[1].(int16))), append(c.Ordinal, int64(r[2].(int)))
			c.Property, c.ValueOrdinal, c.SnakType = append(c.Property, r[3].(string)), append(c.ValueOrdinal, int64(r[4].(int))), append(c.SnakType, r[5].(string))
			c.Datatype, c.ValueType, c.Value = append(c.Datatype, nullText(r[6])), append(c.ValueType, nullText(r[7])), append(c.Value, nullText(r[8]))
			c.ValueEntity, c.ValueJSON = append(c.ValueEntity, nullText(r[9])), append(c.ValueJSON, nullText(r[10]))
		}
	}
	return c
}
