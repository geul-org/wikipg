//ff:type feature=load type=model
//ff:what Config — what AddEntity writes: the dump, the term languages kept and whether the raw line is stored
package load

// Config is what AddEntity writes for each entity: the dump it belongs to,
// the languages of labels, descriptions and aliases to keep (nil or empty:
// all) and whether the raw line goes into entity_raw.
type Config struct {
	DumpID int16
	Langs  map[string]bool
	Raw    bool
}
