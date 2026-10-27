//ff:type feature=source type=model
//ff:what Store — where a loaded dump lives: the schema of sql/wikidata.sql and the dump_id within it
package source

// Store names one loaded dump (sql/wikidata.sql): the schema that holds the
// tables (wikiload -schema) and the dump_id of the dump within it (wikiload
// -dump-id). Every query reads that dump only.
type Store struct {
	Schema string
	Dump   int16
}
