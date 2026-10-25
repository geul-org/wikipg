//ff:func feature=cli type=command control=sequence
//ff:what portOne — load, build and publish one entity's document; false when the entity was never loaded
package main

import (
	"database/sql"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"

	"github.com/geul-org/geul-dev/wiki/go/internal/port"
	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// portOne ports entity id: loads it (source.LoadEntity), builds its document
// under issuer (port.BuildDoc) and publishes it in dir (resolve.Dir.Publish:
// the next version, its h/ copy and the version list). It is false when the
// entity is not in the database.
func portOne(db *sql.DB, dir resolve.Dir, id string, issuer []core.Word, c config) (port.Report, bool, error) {
	e, ok, err := source.LoadEntity(db, id, c.Langs)
	if err != nil || !ok {
		return port.Report{ID: id}, false, err
	}
	b, rep, err := port.BuildDoc(e, issuer, provenanceOf(c.Commit))
	if err != nil {
		return rep, false, err
	}
	if _, err := dir.Publish(b); err != nil {
		return rep, false, err
	}
	return rep, true, nil
}
