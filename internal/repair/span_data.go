//ff:type feature=repair type=model
//ff:what spanData — what a batch's dump lines give: the lines and entities read, reference occurrences, statement_reference rows, bytes, errors
package repair

// spanData is a batch read from entity_raw and parsed: the lines read in
// order and the entity on each, every reference occurrence in dump order,
// the statement_reference rows the loader builds from them, the raw bytes
// read and the lines that did not parse (one message each).
type spanData struct {
	Lines  []int64
	IDs    map[int64]string
	Occs   []occ
	Refs   []stmtRef
	Bytes  int64
	Errors []string
}
