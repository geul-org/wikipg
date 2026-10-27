//ff:func feature=port type=encoder control=sequence
//ff:what monolingualRecord — a Wikidata monolingualtext value as UTF8-LANGUAGE-TEXT; otherwise the reason it is not converted
package port

import (
	"encoding/json"

	"github.com/geul-org/gogeul/pkg/core"
)

// monolingualRecord converts a monolingualtext datavalue
// ({"text":…,"language":…}) to UTF8-LANGUAGE-TEXT (languageText). JSON that
// does not decode and a language code that is not a BCP 47 tag in form give
// their reason. encoding/json turns invalid UTF-8 in the text into U+FFFD, so
// the text is always valid UTF-8 (the raw dump line keeps the original).
func monolingualRecord(value string) (core.Record, string) {
	var v struct {
		Text     string `json:"text"`
		Language string `json:"language"`
	}
	if err := json.Unmarshal([]byte(value), &v); err != nil {
		return core.Record{}, "monolingualtext: JSON 아님"
	}
	r, ok := languageText(v.Language, v.Text)
	if !ok {
		return core.Record{}, "monolingualtext: 언어 코드 " + v.Language
	}
	return r, ""
}
