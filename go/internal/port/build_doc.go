//ff:func feature=port type=encoder control=sequence
//ff:what BuildDoc — the GEUL document of one Wikidata entity under issuer, with its report
package port

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

// BuildDoc writes the document of entity e under issuer (wiki/README): the
// HEAD provenance (headRecords), then LABEL per label, ALIAS per alias and
// STATEMENT per convertible claim, Edge Ids from 1. The document ID is the
// number of e's ID. The report counts what was written and skipped.
func BuildDoc(e source.Entity, issuer []core.Word, p Provenance) ([]byte, Report, error) {
	rep := Report{ID: e.ID, Skipped: map[string]int{}, Raw: map[string]int{}}
	_, num, okNum := entityNumber(e.ID)
	subject, okRef := entityRef(e.ID)
	if !okNum || !okRef {
		return nil, rep, fmt.Errorf("wikiport: %q is not an item or property ID", e.ID)
	}
	head, err := headRecords(issuer, num, subject, p)
	if err != nil {
		return nil, rep, err
	}
	next := uint64(1)
	body := textEdges(edgeLabel, &next, subject, e.Labels)
	body = append(body, textEdges(edgeAlias, &next, subject, e.Aliases)...)
	stmts := statementEdges(&next, subject, e.Claims, &rep)
	rep.Labels, rep.Aliases, rep.Statements = len(e.Labels), len(e.Aliases), len(stmts)
	return document(head, append(body, stmts...)), rep, nil
}
