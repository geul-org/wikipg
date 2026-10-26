//ff:type feature=load type=model
//ff:what DumpInfo — the dump row: which file, its hashes and the loader that reads it
package load

import "time"

// DumpInfo is the dump row (wikidata.sql dump): the dump date and file, where
// it came from, its size and hashes, when it was published and downloaded,
// and the loader and its commit.
type DumpInfo struct {
	ID           int16
	Date         time.Time
	FileName     string
	URL          string
	Size         int64
	MD5          string
	SHA1         string
	SHA256       string
	PublishedAt  time.Time
	DownloadedAt time.Time
	Loader       string
	LoaderCommit string
}
