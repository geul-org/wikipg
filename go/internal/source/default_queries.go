//ff:func feature=source type=loader control=sequence
//ff:what defaultQueries — the geuldev.public SQL for one entity: labels and descriptions by language, aliases in alias_order, claims in triple ID order, qualifiers in load order
package source

// defaultQueries reads geuldev.public: labels and descriptions by language, aliases by
// language then alias_order, claims in triple ID order and qualifiers by
// triple ID then physical row order (ctid). triple_qualifiers has no order
// column; the row order is the loader's insertion order, stable while the
// table is not rewritten. The loader keeps each property's values in dump
// order but walks the properties as a Go map, so the order across properties
// is not Wikidata's qualifiers-order (wiki/README 빈칸). The claim IDs are
// passed as an array rather than joined: the join is estimated so high that
// PostgreSQL JIT-compiles it, about 90 ms per entity instead of 0.5 ms.
func defaultQueries() entityQueries {
	return entityQueries{
		Exists:       "SELECT count(*) FROM entities WHERE id = $1",
		Labels:       "SELECT language, label FROM entity_labels WHERE entity_id = $1 AND language = ANY($2) ORDER BY language",
		Descriptions: "SELECT language, description FROM entity_descriptions WHERE entity_id = $1 AND language = ANY($2) AND description IS NOT NULL ORDER BY language",
		Aliases:      "SELECT language, alias FROM entity_aliases WHERE entity_id = $1 AND language = ANY($2) ORDER BY language, alias_order",
		Claims:       "SELECT id, property, object_type, coalesce(object_value, ''), coalesce(rank, '') FROM triples WHERE subject = $1 ORDER BY id",
		Qualifiers: "SELECT triple_id, property, coalesce(datatype, ''), coalesce(value, '') FROM triple_qualifiers " +
			"WHERE triple_id = ANY(ARRAY(SELECT id FROM triples WHERE subject = $1)) ORDER BY triple_id, ctid",
	}
}
