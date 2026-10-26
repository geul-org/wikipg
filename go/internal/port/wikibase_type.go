//ff:func feature=port type=helper control=selection
//ff:what wikibaseType — the Wikibase ontology name of a property datatype (url → Url); false for an unknown datatype
package port

// wikibaseType maps a Wikidata property datatype to its class in the Wikibase
// RDF ontology (http://wikiba.se/ontology#…, files/geul/
// wikidata-datatype-mapping.md): the 18 datatypes of Special:ListDatatypes.
func wikibaseType(datatype string) (string, bool) {
	switch datatype {
	case "wikibase-item":
		return "WikibaseItem", true
	case "wikibase-property":
		return "WikibaseProperty", true
	case "wikibase-lexeme":
		return "WikibaseLexeme", true
	case "wikibase-form":
		return "WikibaseForm", true
	case "wikibase-sense":
		return "WikibaseSense", true
	case "entity-schema":
		return "EntitySchema", true
	case "string":
		return "String", true
	case "external-id":
		return "ExternalId", true
	case "url":
		return "Url", true
	case "commonsMedia":
		return "CommonsMedia", true
	case "geo-shape":
		return "GeoShape", true
	case "tabular-data":
		return "TabularData", true
	case "monolingualtext":
		return "Monolingualtext", true
	case "time":
		return "Time", true
	case "globe-coordinate":
		return "GlobeCoordinate", true
	case "quantity":
		return "Quantity", true
	case "math":
		return "Math", true
	case "musical-notation":
		return "MusicalNotation", true
	default:
		return "", false
	}
}
