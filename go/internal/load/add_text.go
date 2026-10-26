//ff:func feature=load type=model control=sequence
//ff:what Rows.addText — adds one text row, recording a load_error when NUL bytes had to be removed
package load

// addText adds a row to table whose last value is text, removing NUL bytes
// (cleanText) and recording a load_error of the line when it had any.
func (r *Rows) addText(c Config, line int64, id, table, text string, values ...any) {
	clean, ok := cleanText(text)
	if !ok {
		r.Error(c.DumpID, line, id, table, "NUL byte removed from "+table+" text", nil)
	}
	r.add(table, append(values, clean)...)
}
