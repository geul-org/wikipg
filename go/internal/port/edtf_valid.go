//ff:func feature=port type=helper control=sequence
//ff:what edtfValid — whether go-edtf parses the EDTF string, real days included; otherwise its error
package port

import "github.com/sfomuseum/go-edtf/parser"

// edtfValid checks s with github.com/sfomuseum/go-edtf (BSD-3-Clause,
// THIRD_PARTY_NOTICES.md). parser.ParseString is used rather than IsValid:
// IsValid only matches the syntax (it accepts 2023-02-29 and 1974-13), while
// ParseString also checks the days of the month. go-edtf does not support
// Y-prefixed years yet and reports them as an error.
func edtfValid(s string) error {
	_, err := parser.ParseString(s)
	return err
}
