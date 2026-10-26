//ff:func feature=load type=model control=sequence
//ff:what tables — the wikidata tables in write order with their COPY columns (wiki/wikidata.sql)
package load

// tables lists the tables of wiki/wikidata.sql the loader writes, in the
// order a batch writes them, with their columns (dump is written once by
// InsertDump, not per batch).
func tables() []table {
	snak := []string{"snaktype", "datatype", "value_type", "value", "value_entity"}
	return []table{
		{"entity_raw", []string{"entity_id", "dump_id", "line_no", "raw"}},
		{"entity", []string{"entity_id", "line_no", "kind", "num", "type", "datatype", "lastrevid", "modified", "dump_id"}},
		{"label", []string{"entity_id", "dump_id", "language", "value"}},
		{"description", []string{"entity_id", "dump_id", "language", "value"}},
		{"alias", []string{"entity_id", "dump_id", "language", "ordinal", "value"}},
		{"sitelink", []string{"entity_id", "dump_id", "site", "title", "badges"}},
		{"statement", append([]string{"statement_id", "guid", "entity_id", "property", "ordinal", "rank"}, append(snak, "dump_id")...)},
		{"qualifier", append([]string{"statement_id", "ordinal", "property", "value_ordinal"}, append(snak, "hash")...)},
		{"reference", []string{"hash", "dump_id"}},
		{"statement_reference", []string{"statement_id", "ordinal", "hash"}},
		{"reference_snak", append([]string{"hash", "dump_id", "ordinal", "property", "value_ordinal"}, snak...)},
		{"lexeme", []string{"entity_id", "dump_id", "language", "lexical_category"}},
		{"lemma", []string{"entity_id", "dump_id", "language", "value"}},
		{"form", []string{"entity_id", "lexeme_id", "dump_id", "ordinal", "grammatical_features"}},
		{"form_representation", []string{"entity_id", "dump_id", "language", "value"}},
		{"sense", []string{"entity_id", "lexeme_id", "dump_id", "ordinal"}},
		{"sense_gloss", []string{"entity_id", "dump_id", "language", "value"}},
		{"load_error", []string{"dump_id", "line_no", "entity_id", "stage", "message", "excerpt"}},
	}
}
