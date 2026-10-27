//ff:func feature=port type=encoder control=sequence
//ff:what TestDocument — HEAD, BODY marker, BODY records and END inside DOC
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocument(t *testing.T) {
	b := document([]core.Record{core.EncodeSpecVersion(core.SpecVersion)}, []core.Record{core.EncodeUintValue(1)})
	ws, _ := core.ReadWords(b)
	d, err := core.SplitDocument(ws)
	if err != nil || len(d.Head) != 1 || len(d.Body) != 1 {
		t.Fatalf("got %+v %v", d, err)
	}
}
