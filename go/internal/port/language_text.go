//ff:func feature=port type=encoder control=sequence
//ff:what languageText — a UTF8-LANGUAGE-TEXT Value from a language code and text; false when the code is not a BCP 47 tag in form
package port

import (
	"regexp"

	"github.com/geul-org/geul/go/pkg/core"
)

var bcp47Form = regexp.MustCompile(`^[A-Za-z]{2,8}(-[A-Za-z0-9]{1,8})*$`)

// languageText builds the UTF8-LANGUAGE-TEXT of text in language lang, the
// language as a LANGUAGE-BCP47-CODE (the GEUL language registry for
// LANGUAGE-SHORT-CODE is not built yet). lang must have the form of a BCP 47
// tag (letters, then hyphen-separated subtags); only the form is checked, so a
// Wikimedia code of that form that BCP 47 does not register (simple) passes.
func languageText(lang, text string) (core.Record, bool) {
	if !bcp47Form.MatchString(lang) {
		return core.Record{}, false
	}
	return core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(lang)), []byte(text)), true
}
