//ff:func feature=load type=model control=sequence
//ff:what Rows.addText — adds one text row: the text without NUL bytes and, when it had any, its exact JSON literal
package load

// addText adds a row to table whose last two values are the text with NUL
// bytes removed (cleanText, text cannot hold them) and value_json, the exact
// value as a JSON string literal when it had any (nulJSON). Nothing is lost.
func (r *Rows) addText(table, text string, values ...any) {
	clean, _ := cleanText(text)
	r.add(table, append(values, clean, nulJSON(text))...)
}
