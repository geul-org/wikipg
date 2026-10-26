//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestObjectRecord — entity IDs, strings, dates and monolingual texts convert; an unconvertible date has its time reason; empty, lexeme, invalid UTF-8 and unsupported values are skipped
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"

	"github.com/geul-org/geul-dev/wiki/go/internal/source"
)

func TestObjectRecord(t *testing.T) {
	enc := func(c source.Claim) []core.Word {
		r, reason := objectRecord(c)
		if reason != "" {
			t.Fatalf("%+v: %q", c, reason)
		}
		ws, _ := core.EncodeRecord(r)
		return ws
	}
	if ws := enc(source.Claim{Type: "wikibase-entityid", Value: "Q5"}); !slices.Equal(ws, []core.Word{0x00A0, 5}) {
		t.Fatalf("item: %04X", ws)
	}
	if ws := enc(source.Claim{Type: "string", Value: "ab"}); !slices.Equal(ws, []core.Word{0x08C0, 0x6162}) {
		t.Fatalf("string: %04X", ws)
	}
	if ws := enc(source.Claim{Type: "time", Value: `{"time":"+1970-01-01T00:00:00Z","precision":11,"calendarmodel":"` + gregorian + `"}`}); !slices.Equal(ws, []core.Word{0x0920, 0, 0}) {
		t.Fatalf("time: %04X", ws)
	}
	if ws := enc(source.Claim{Type: "monolingualtext", Value: `{"text":"Paris","language":"fr"}`}); !slices.Equal(ws, []core.Word{0x0ACB, 0x08C6, 0x6672, 0x5061, 0x7269, 0x7300}) {
		t.Fatalf("monolingual: %04X", ws)
	}
	for want, c := range map[string]source.Claim{
		"값 없음(novalue·somevalue)":    {Type: "wikibase-item"},
		"개체: L":                      {Type: "wikibase-entityid", Value: "L16"},
		"문자열: UTF-8 아님":              {Type: "external-id", Value: "\xff"},
		"형식 미지원: quantity":           {Type: "quantity", Value: "{}"},
		"monolingualtext: 언어 코드 x y": {Type: "monolingualtext", Value: `{"text":"a","language":"x y"}`},
		"time: JSON 아님":              {Type: "time", Value: "x"},
	} {
		if _, got := objectRecord(c); got != want {
			t.Fatalf("%+v: %q want %q", c, got, want)
		}
	}
}
