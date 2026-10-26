//ff:func feature=source type=loader control=sequence
//ff:what LoadEntity — an entity's labels, descriptions, aliases and claims from a loaded dump; false when the ID is not in it
package source

import (
	"database/sql"
	"fmt"
)

// LoadEntity reads the entity id from the dump s (queriesFor): its labels,
// descriptions and aliases in langs and all its claims. It is false, with no
// error, when the dump has no such ID. An invalid schema name is an error.
func LoadEntity(db *sql.DB, s Store, id string, langs []string) (Entity, bool, error) {
	if !validSchema(s.Schema) {
		return Entity{}, false, fmt.Errorf("wikiport: schema %q is not a plain identifier", s.Schema)
	}
	return loadEntityWith(db, queriesFor(s), id, langs)
}
