//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestLanguageText — a Wikimedia code is mapped; a tag-shaped code gives the docs' "Paris"@fr words; malformed codes are false
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLanguageText(t *testing.T) {
	r, ok := languageText("fr", "Paris")
	ws, _ := core.EncodeRecord(r)
	if !ok || !slices.Equal(ws, []core.Word{0x0ACB, 0x08C6, 0x6672, 0x5061, 0x7269, 0x7300}) {
		t.Fatalf("got %04X %v", ws, ok)
	}
	r, _ = languageText("sr-el", "Beograd")
	lang, _, err := core.DecodeLanguageText(r)
	if b, _ := core.DecodeBytesValue(lang, core.CodecLanguageBCP47Even); err != nil || string(b) != "sr-Latn" {
		t.Fatalf("sr-el: %q %v", b, err)
	}
	if _, ok := languageText("zh-Hant", "x"); !ok {
		t.Fatal("zh-Hant: want ok")
	}
	for _, bad := range []string{"", "e", "en_US", "en-", "toolongcode", "ko KR"} {
		if _, ok := languageText(bad, "x"); ok {
			t.Fatalf("%q: want false", bad)
		}
	}
}
