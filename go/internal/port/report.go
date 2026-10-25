//ff:type feature=port type=model
//ff:what Report — what one ported document holds and what it left out, with the properties it used
package port

// Report tells what one ported document holds and what it left out: counts
// of LABEL, ALIAS and STATEMENT records, skipped claims by reason (unsupported
// value type, novalue/somevalue, date precision …) and the properties the
// document uses as predicates or values, for porting those next.
type Report struct {
	ID         string
	Labels     int
	Aliases    int
	Statements int
	Skipped    map[string]int
	Properties []string
}
