//ff:func feature=port type=encoder control=selection
//ff:what datatypeEdges — a property's datatype in its own document: REFERENCE-NODE(#n → wikibase propertyType URI) and TRIPLE(#n, P, the type's URI)
package port

import "github.com/geul-org/geul/go/pkg/core"

// wikibaseOntology is the Wikibase RDF ontology namespace.
const wikibaseOntology = "http://wikiba.se/ontology#"

// datatypeEdges states the datatype of the property subject inside its own
// document (wiki/README P 문서), so a reader that meets the property resolves
// one document: REFERENCE-NODE binds a LOCAL predicate to the URI of
// wikibase:propertyType, and TRIPLE(predicate, subject, URI of the type,
// normal) gives the type (wikibaseType). No datatype (an item) gives nothing;
// an unknown one is counted in rep.Skipped and gives nothing. Edge Ids and
// the LOCAL come from *next.
func datatypeEdges(a IDs, subject core.Record, datatype string, rep *Report) []core.Record {
	uri := func(s string) core.Record { return core.EncodeBytesValue(codecURIEven, []byte(s)) }
	name, ok := wikibaseType(datatype)
	switch {
	case datatype == "":
		return nil
	case !ok:
		rep.Skipped["데이터 타입: "+datatype]++
		return nil
	}
	property := uri(wikibaseOntology + "propertyType")
	bind, pred := a.Define(property)
	triple, _ := tripleEdge(a, core.EncodeLocalRef(pred), subject, uri(wikibaseOntology+name), 0)
	return []core.Record{compose(core.KindEdge, edgeReferenceNode, core.EncodeLocalRef(bind), core.EncodeLocalRef(pred), property), triple}
}
