//ff:func feature=source type=loader control=sequence
//ff:what LoadEntity — an entity's labels, aliases and claims from geuldev; false when the ID was never loaded
package source

import "database/sql"

// LoadEntity reads the entity id from geuldev.public (defaultQueries): its
// labels and aliases in langs and all its claims. It is false, with no error,
// when entities has no such ID.
func LoadEntity(db *sql.DB, id string, langs []string) (Entity, bool, error) {
	return loadEntityWith(db, defaultQueries(), id, langs)
}
