//ff:func feature=cli type=command control=sequence
//ff:what run — port the QID range, then the properties they use, write the name index and the report
package main

import (
	"database/sql"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
	"github.com/geul-org/wikiport/go/internal/source"
	_ "github.com/lib/pq"
)

// run reads the loaded dump's provenance (source.LoadDump, provenanceOf) and
// ports Q<from>…Q<to>, then every property those documents use, all under
// the one Issuer wikidata.geul.org (0x00010000; document IDs by
// port.DocNumber), writes its documents.geul and, when asked, the report.
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
	d, err := source.LoadDump(db, source.Store{Schema: c.Schema, Dump: c.Dump})
	if err != nil {
		return fmt.Errorf("dump %s/%d: %w", c.Schema, c.Dump, err)
	}
	prov, err := provenanceOf(d, c.Commit, c.LibCommit)
	if err != nil {
		return err
	}
	dir := resolve.Dir{Root: c.Root}
	issuer := []core.Word{1, 0}
	qReps, qDocs, err := portIDs(db, dir, "Q", rangeNumbers(c.From, c.To), issuer, c, prov)
	if err != nil {
		return err
	}
	pReps, pDocs, err := portIDs(db, dir, "P", propertyNumbers(qReps), issuer, c, prov)
	if err != nil {
		return err
	}
	if err := writeIndex(c.Root, issuer, append(qDocs, pDocs...)); err != nil {
		return err
	}
	fmt.Printf("Q: %d documents of %d IDs, P: %d documents of %d IDs\n", len(qDocs), len(qReps), len(pDocs), len(pReps))
	if c.Report == "" {
		return nil
	}
	return writeReport(c.Report, append(qReps, pReps...))
}
