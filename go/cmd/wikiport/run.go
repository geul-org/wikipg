//ff:func feature=cli type=command control=sequence
//ff:what run — port the QID range, then the properties they use, write both name indexes and the report
package main

import (
	"database/sql"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
	_ "github.com/lib/pq"
)

// run ports Q<from>…Q<to> under wikidata-qid.geul.org (0x00010000), then every
// property those documents use under wikidata-pid.geul.org (0x00010001),
// writes both documents.geul and, when asked, the report.
func run(args []string) error {
	c, err := parseFlags(args)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", c.DSN)
	if err != nil {
		return err
	}
	defer db.Close()
	dir := resolve.Dir{Root: c.Root}
	qIssuer, pIssuer := []core.Word{1, 0}, []core.Word{1, 1}
	qReps, qDocs, err := portIDs(db, dir, "Q", rangeNumbers(c.From, c.To), qIssuer, c)
	if err != nil {
		return err
	}
	pReps, pDocs, err := portIDs(db, dir, "P", propertyNumbers(qReps), pIssuer, c)
	if err != nil {
		return err
	}
	if err := writeIndex(c.Root, qIssuer, "q", qDocs); err != nil {
		return err
	}
	if err := writeIndex(c.Root, pIssuer, "p", pDocs); err != nil {
		return err
	}
	fmt.Printf("Q: %d documents of %d IDs, P: %d documents of %d IDs\n", len(qDocs), len(qReps), len(pDocs), len(pReps))
	if c.Report == "" {
		return nil
	}
	return writeReport(c.Report, append(qReps, pReps...))
}
