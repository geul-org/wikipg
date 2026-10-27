//ff:func feature=source type=loader control=sequence
//ff:what LoadDump — the provenance row of a loaded dump; an error for a bad schema name or a missing dump
package source

import (
	"database/sql"
	"fmt"
)

// LoadDump reads the dump row of s (URL, file SHA-256, download time, loader
// and its commit). An invalid schema name or a dump_id with no row is an error.
func LoadDump(db *sql.DB, s Store) (Dump, error) {
	if !validSchema(s.Schema) {
		return Dump{}, fmt.Errorf("wikiport: schema %q is not a plain identifier", s.Schema)
	}
	var d Dump
	err := db.QueryRow(fmt.Sprintf("SELECT url, sha256, downloaded_at, loader, loader_commit FROM %s.dump WHERE dump_id = $1", s.Schema), s.Dump).
		Scan(&d.URL, &d.SHA256, &d.DownloadedAt, &d.Loader, &d.LoaderCommit)
	return d, err
}
