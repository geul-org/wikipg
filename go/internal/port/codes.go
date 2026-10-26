//ff:type feature=port type=model
//ff:what codes — the GEUL codes this porter writes that the library does not name, and the Wikidata constants it reads
package port

// Codes the GEUL library (core) has no constant for: edge/ALIAS,
// edge/DESCRIPTION, edge/TRIPLE,
// edge/QUALIFIER, the Meta this porter writes in HEAD and EXTERNAL-ISSUER-1WORD.
const (
	codecURIEven      uint8 = 0x44 // Value URI (even bytes)
	edgeLabel         uint8 = 0x04 // Edge LABEL
	edgeAlias         uint8 = 0x05 // Edge ALIAS (edge/ALIAS)
	edgeDescription   uint8 = 0x06 // Edge DESCRIPTION (edge/DESCRIPTION)
	edgeTriple        uint8 = 0x0A // Edge TRIPLE (edge/TRIPLE)
	edgeQualifier     uint8 = 0x0B // Edge QUALIFIER (edge/QUALIFIER)
	metaSource        uint8 = 0x26 // Meta SOURCE
	metaAcquiredAt    uint8 = 0x28 // Meta ACQUIRED-AT
	metaTransformedBy uint8 = 0x33 // Meta TRANSFORMED-BY

	refExternalIssuer1Word uint8 = 0x0C // Reference CCCC EXTERNAL-ISSUER-1WORD (reference/EXTERNAL-ISSUER)

	gregorian = "http://www.wikidata.org/entity/Q1985727" // Wikidata calendar model: proleptic Gregorian
	julian    = "http://www.wikidata.org/entity/Q1985786" // Wikidata calendar model: proleptic Julian
)
