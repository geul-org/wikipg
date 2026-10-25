//ff:type feature=port type=model
//ff:what Tool — a program in the provenance chain: its repository URI and the commit or version that ran
package port

// Tool is one program in the provenance chain, written as TRANSFORMED-BY: its
// repository URI and the commit (or version text) that ran.
type Tool struct {
	URI     string
	Version string
}
