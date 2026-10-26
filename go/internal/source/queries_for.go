//ff:func feature=source type=loader control=sequence
//ff:what queriesFor — the SQL for one entity of a loaded dump: existence, texts by language, claims and qualifiers in dump order
package source

import "fmt"

// queriesFor reads one entity of dump s.Dump from schema s.Schema (wiki/
// wikidata.sql; the name must pass validSchema): labels and descriptions by
// language, aliases by language then dump order, statements by statement_id
// (dump order) and qualifiers by statement then their dump order
// (qualifiers-order). A value is the bare entity ID for an entity
// (value_entity), the unquoted string for a JSON string, else the datavalue
// JSON as text (jsonb's spacing and key order); novalue and somevalue give an
// empty value. The statement IDs are passed as an array rather than joined,
// which PostgreSQL would plan as a slow JIT-compiled join.
func queriesFor(s Store) entityQueries {
	value := "coalesce(value_entity, value #>> '{}', '')"
	return entityQueries{
		Exists:       fmt.Sprintf("SELECT count(*) FROM %s.entity WHERE entity_id = $1 AND dump_id = %d", s.Schema, s.Dump),
		Labels:       fmt.Sprintf("SELECT language, value FROM %s.label WHERE entity_id = $1 AND dump_id = %d AND language = ANY($2) ORDER BY language", s.Schema, s.Dump),
		Descriptions: fmt.Sprintf("SELECT language, value FROM %s.description WHERE entity_id = $1 AND dump_id = %d AND language = ANY($2) ORDER BY language", s.Schema, s.Dump),
		Aliases:      fmt.Sprintf("SELECT language, value FROM %s.alias WHERE entity_id = $1 AND dump_id = %d AND language = ANY($2) ORDER BY language, ordinal", s.Schema, s.Dump),
		Claims: fmt.Sprintf("SELECT statement_id, property, coalesce(value_type, ''), coalesce(datatype, ''), %s, rank::text FROM %s.statement "+
			"WHERE entity_id = $1 AND dump_id = %d ORDER BY statement_id", value, s.Schema, s.Dump),
		Qualifiers: fmt.Sprintf("SELECT statement_id, property, coalesce(value_type, ''), coalesce(datatype, ''), %s FROM %s.qualifier "+
			"WHERE statement_id = ANY(ARRAY(SELECT statement_id FROM %s.statement WHERE entity_id = $1 AND dump_id = %d)) ORDER BY statement_id, ordinal",
			value, s.Schema, s.Schema, s.Dump),
	}
}
