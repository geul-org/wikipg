//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestCommonsURI — a file name becomes its File: page with _ for spaces; a Data page and a title with ? stay one segment
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCommonsURI(t *testing.T) {
	for want, r := range map[string]core.Record{
		"https://commons.wikimedia.org/wiki/File:Douglas_adams_portrait.jpg":     commonsURI("File:", "Douglas adams portrait.jpg"),
		"https://commons.wikimedia.org/wiki/Data:Germany.map":                    commonsURI("", "Data:Germany.map"),
		"https://commons.wikimedia.org/wiki/File:What%3F_%EC%82%AC%EC%A7%84.jpg": commonsURI("File:", "What? 사진.jpg"),
	} {
		if b, err := core.DecodeBytesValue(r, codecURIEven); err != nil || string(b) != want {
			t.Fatalf("got %q want %q", b, want)
		}
	}
}
