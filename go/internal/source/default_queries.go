//ff:func feature=source type=loader control=sequence
//ff:what defaultQueries — the geuldev.public SQL for one entity: labels by language, aliases in alias_order, claims in triple ID order
package source

// defaultQueries reads geuldev.public: labels by language, aliases by
// language then alias_order, and claims in triple ID order, so a rerun lays
// them out the same way.
func defaultQueries() entityQueries {
	return entityQueries{
		Exists:  "SELECT count(*) FROM entities WHERE id = $1",
		Labels:  "SELECT language, label FROM entity_labels WHERE entity_id = $1 AND language = ANY($2) ORDER BY language",
		Aliases: "SELECT language, alias FROM entity_aliases WHERE entity_id = $1 AND language = ANY($2) ORDER BY language, alias_order",
		Claims:  "SELECT id, property, object_type, coalesce(object_value, ''), coalesce(rank, '') FROM triples WHERE subject = $1 ORDER BY id",
	}
}
