//ff:func feature=port type=encoder control=sequence
//ff:what BuildDoc — the GEUL document of one Wikidata entity under issuer, with its report
package port

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/wikiport/go/internal/source"
)

// BuildDoc writes the document of entity e under issuer (wiki/README): the
// HEAD provenance (headRecords), then LABEL per label, DESCRIPTION per
// description, ALIAS per alias, a property's datatype (datatypeEdges) and a
// TRIPLE followed by its QUALIFIERs per convertible claim. Edge Ids come from a:
// a statement equal to one of the previous version keeps its ID, new ones
// take the next (OVERVIEW 개정, delta.LocalIDs). The document ID is the
// DocNumber of e's ID (its letter and number). The report counts what was written and skipped.
func BuildDoc(e source.Entity, issuer []core.Word, p Provenance, a IDs) ([]byte, Report, error) {
	rep := Report{ID: e.ID, Skipped: map[string]int{}, Raw: map[string]int{}}
	prefix, n, okNum := entityNumber(e.ID)
	subject, okRef := entityRef(e.ID)
	if !okNum || !okRef {
		return nil, rep, fmt.Errorf("wikiport: %q is not an item or property ID", e.ID)
	}
	num, err := DocNumber(prefix, n)
	if err != nil {
		return nil, rep, err
	}
	head, err := headRecords(issuer, num, subject, p)
	if err != nil {
		return nil, rep, err
	}
	labels := textEdges(edgeLabel, a, subject, e.Labels, &rep)
	descriptions := textEdges(edgeDescription, a, subject, e.Descriptions, &rep)
	aliases := textEdges(edgeAlias, a, subject, e.Aliases, &rep)
	body := append(append(labels, descriptions...), aliases...)
	body = append(body, datatypeEdges(a, subject, e.Datatype, &rep)...)
	body = append(body, tripleEdges(a, subject, e.Claims, &rep)...)
	rep.Labels, rep.Descriptions, rep.Aliases = len(labels), len(descriptions), len(aliases)
	return document(head, body), rep, nil
}
