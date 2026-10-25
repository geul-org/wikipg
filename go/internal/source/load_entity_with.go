//ff:func feature=source type=loader control=sequence
//ff:what loadEntityWith — an entity's labels, aliases and claims by the given queries; false when the ID was never loaded
package source

import (
	"database/sql"

	"github.com/lib/pq"
)

// loadEntityWith reads the entity id by q: existence, then labels and aliases
// in langs, then claims. It is false, with no error, when the entity does not
// exist.
func loadEntityWith(db *sql.DB, q entityQueries, id string, langs []string) (Entity, bool, error) {
	ok, err := entityExists(db, q.Exists, id)
	if err != nil || !ok {
		return Entity{}, false, err
	}
	e := Entity{ID: id}
	if e.Labels, err = loadTexts(db, q.Labels, id, pq.Array(langs)); err != nil {
		return Entity{}, false, err
	}
	if e.Aliases, err = loadTexts(db, q.Aliases, id, pq.Array(langs)); err != nil {
		return Entity{}, false, err
	}
	if e.Claims, err = loadClaims(db, q.Claims, id); err != nil {
		return Entity{}, false, err
	}
	return e, true, nil
}
