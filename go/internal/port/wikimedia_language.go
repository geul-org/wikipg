//ff:func feature=port type=helper control=sequence
//ff:what wikimediaLanguage — the BCP 47 tag of a Wikimedia language code that is not one (files/geul/wikimedia-language-mapping.md); others as given
package port

import "strings"

// wikimediaBCP47 maps the Wikimedia language codes Wikidata uses that are not
// BCP 47 tags to the tag they stand for (2026-10-26 full scan: these 14 are
// every invalid code in labels, aliases, descriptions and monolingual text).
// Each tag rests on its IANA Language Subtag Registry entry; the private use
// tags (x-formal, x-informal, x-bms, x-tara) are GEUL definitions.
var wikimediaBCP47 = map[string]string{
	"sr-el": "sr-Latn", "bat-smg": "sgs", "fiu-vro": "vro", "cbk-zam": "cbk", "simple": "en-simple",
	"eml": "egl", "roa-rup": "rup", "zh-classical": "lzh",
	"map-bms": "jv-x-bms", "roa-tara": "nap-x-tara",
	"de-formal": "de-x-formal", "es-formal": "es-x-formal", "hu-formal": "hu-x-formal", "nl-informal": "nl-x-informal",
}

// wikimediaLanguage returns the BCP 47 tag for the Wikimedia code lang
// (case-insensitive), or lang unchanged when it is not in wikimediaBCP47.
func wikimediaLanguage(lang string) string {
	if tag, ok := wikimediaBCP47[strings.ToLower(lang)]; ok {
		return tag
	}
	return lang
}
