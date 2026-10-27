//ff:func feature=port type=encoder control=sequence
//ff:what languageText — a UTF8-LANGUAGE-TEXT Value from a language code (Wikimedia codes mapped) and text; false when the code is not a BCP 47 tag in form
package port

import (
	"regexp"

	"github.com/geul-org/gogeul/pkg/core"
)

var bcp47Form = regexp.MustCompile(`^[A-Za-z]{2,8}(-[A-Za-z0-9]{1,8})*$`)

// languageText builds the UTF8-LANGUAGE-TEXT of text in language lang, the
// language as a LANGUAGE-BCP47-CODE (the GEUL language registry for
// LANGUAGE-SHORT-CODE is not built yet). A Wikimedia code that is not BCP 47
// is first mapped to its tag (wikimediaLanguage). The tag must then have the
// form of a BCP 47 tag (letters, then hyphen-separated subtags); only the form
// is checked.
func languageText(lang, text string) (core.Record, bool) {
	lang = wikimediaLanguage(lang)
	if !bcp47Form.MatchString(lang) {
		return core.Record{}, false
	}
	return core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(lang)), []byte(text)), true
}
