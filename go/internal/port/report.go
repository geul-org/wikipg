//ff:type feature=port type=model
//ff:what Report — what one ported document holds and what it left out, with the properties it used
package port

// Report tells what one ported document holds and what it left out: counts
// of LABEL, DESCRIPTION, ALIAS, TRIPLE and QUALIFIER records, skipped claims and
// qualifiers by reason (unsupported value type, novalue/somevalue …), dates
// kept as original text by conversion reason (Raw) and the properties the
// document uses as predicates, qualifier properties or values, for porting
// those next.
type Report struct {
	ID           string
	Labels       int
	Descriptions int
	Aliases      int
	Triples      int
	Qualifiers   int
	Skipped      map[string]int
	Raw          map[string]int
	Properties   []string
}
