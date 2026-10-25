//ff:type feature=port type=model
//ff:what codes — the GEUL codes this porter writes that the library does not name, and the Wikidata constants it reads
package port

// Codes the GEUL library (core) has no constant for. The STATEMENT layout is
// this porter (wiki/README 값 변환); ALIAS is edge/ALIAS.
const (
	codecUTF8NFCTextEven uint8 = 0x42 // Value UTF8-NFC-TEXT (even bytes)
	codecURIEven         uint8 = 0x44 // Value URI (even bytes)
	edgeLabel            uint8 = 0x04 // Edge LABEL
	edgeAlias            uint8 = 0x05 // Edge ALIAS (edge/ALIAS)
	edgeStatement        uint8 = 0x0B // Edge STATEMENT (proposal layout)
	metaSource           uint8 = 0x26 // Meta SOURCE
	metaAcquiredAt       uint8 = 0x28 // Meta ACQUIRED-AT
	metaTransformedBy    uint8 = 0x33 // Meta TRANSFORMED-BY

	gregorian = "http://www.wikidata.org/entity/Q1985727" // Wikidata calendar model: proleptic Gregorian
	julian    = "http://www.wikidata.org/entity/Q1985786" // Wikidata calendar model: proleptic Julian
)
