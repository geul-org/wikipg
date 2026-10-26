//ff:func feature=load type=writer control=sequence
//ff:what InsertDump — writes the dump row once and marks the load started; an existing row is kept
package load

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// InsertDump writes d as the dump row of schema with load_started now. A row
// with the same dump_id is left as it is (a resumed load keeps its first
// record).
func InsertDump(db *sql.DB, schema string, d DumpInfo) error {
	var published any
	if !d.PublishedAt.IsZero() {
		published = d.PublishedAt
	}
	_, err := db.Exec("INSERT INTO "+pq.QuoteIdentifier(schema)+".dump "+
		"(dump_id, dump_date, file_name, url, size_bytes, md5, sha1, sha256, published_at, downloaded_at, loader, loader_commit, load_started) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (dump_id) DO NOTHING",
		d.ID, d.Date, d.FileName, d.URL, d.Size, nullable(d.MD5), nullable(d.SHA1), d.SHA256, published, d.DownloadedAt, d.Loader, d.LoaderCommit, time.Now())
	return err
}
