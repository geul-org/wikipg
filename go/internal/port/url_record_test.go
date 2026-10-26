//ff:func feature=port type=encoder control=sequence
//ff:what TestURLRecord — an ASCII URL stays, a Korean IRI is percent-encoded into a URI Value, no scheme gives a reason
package port

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestURLRecord(t *testing.T) {
	r, reason := urlRecord("https://douglasadams.com/")
	if b, err := core.DecodeBytesValue(r, codecURIEven); reason != "" || err != nil || string(b) != "https://douglasadams.com/" {
		t.Fatalf("ascii: %q %v %q", b, err, reason)
	}
	r, _ = urlRecord("https://ko.wikipedia.org/wiki/더글러스 애덤스")
	if b, _ := core.DecodeBytesValue(r, codecURIEven); string(b) != "https://ko.wikipedia.org/wiki/%EB%8D%94%EA%B8%80%EB%9F%AC%EC%8A%A4%20%EC%95%A0%EB%8D%A4%EC%8A%A4" {
		t.Fatalf("iri: %q", b)
	}
	if _, reason := urlRecord("douglasadams.com"); reason != "url: 스킴 없음" {
		t.Fatalf("no scheme: %q", reason)
	}
}
