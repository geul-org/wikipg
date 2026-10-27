//ff:type feature=port type=model
//ff:what Tool — a program in the provenance chain: its geul.org term definition document and the source commit that ran
package port

// Tool is one program in the provenance chain, written as TRANSFORMED-BY
// (meta/lineage/TRANSFORMED-BY): Doc is the number n of its geul.org
// term definition document 0x0001000000000000+n (ISSUER 용어 정의), Commit the
// full git object ID of the source that ran, in hex (40 digits for SHA-1, 64
// for SHA-256).
type Tool struct {
	Doc    uint16
	Commit string
}
