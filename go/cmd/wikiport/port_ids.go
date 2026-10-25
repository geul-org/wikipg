//ff:func feature=cli type=command control=iteration dimension=1
//ff:what portIDs — port each entity ID in order; the reports and numbers of the documents written
package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"

	"github.com/geul-org/geul-dev/wiki/go/internal/port"
)

// portIDs ports prefix+n for each n under issuer. It returns the reports of
// all IDs (missing ones too, with no counts) and the numbers of the
// documents written. The first error stops it.
func portIDs(db *sql.DB, dir resolve.Dir, prefix string, nums []uint64, issuer []core.Word, c config) ([]port.Report, []uint64, error) {
	var reps []port.Report
	var written []uint64
	for _, n := range nums {
		id := prefix + strconv.FormatUint(n, 10)
		rep, ok, err := portOne(db, dir, id, issuer, c)
		if err != nil {
			return reps, written, fmt.Errorf("%s: %w", id, err)
		}
		reps = append(reps, rep)
		if ok {
			written = append(written, n)
		}
	}
	return reps, written, nil
}
