//ff:func feature=load type=model control=sequence
//ff:what Rows.Error — records a load_error row for a line: stage, message and the first 1 KB of the line
package load

// Error records one load_error row for line (of entity id, "" when unknown)
// at stage with message, keeping the first 1024 bytes of raw as the excerpt
// with NUL bytes removed (cleanText).
func (r *Rows) Error(dumpID int16, line int64, id, stage, message string, raw []byte) {
	if len(raw) > 1024 {
		raw = raw[:1024]
	}
	excerpt, _ := cleanText(string(raw))
	var ent any
	if id != "" {
		ent = id
	}
	msg, _ := cleanText(message)
	r.add("load_error", dumpID, line, ent, stage, msg, excerpt)
	r.Errors++
}
