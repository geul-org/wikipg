//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestWikimediaLanguage — mapped codes (standard, private use, any case) become their tags; BCP 47 codes and unknowns stay
package port

import "testing"

func TestWikimediaLanguage(t *testing.T) {
	for in, want := range map[string]string{
		"sr-el": "sr-Latn", "simple": "en-simple", "zh-classical": "lzh", "map-bms": "jv-x-bms", "DE-Formal": "de-x-formal",
		"de": "de", "zh-Hant": "zh-Hant", "": "",
	} {
		if got := wikimediaLanguage(in); got != want {
			t.Fatalf("%q: %q want %q", in, got, want)
		}
	}
}
