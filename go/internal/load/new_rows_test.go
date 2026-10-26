//ff:func feature=load type=model control=sequence
//ff:what TestNewRows — an empty batch with no rows and no errors
package load

import "testing"

func TestNewRows(t *testing.T) {
	r := NewRows()
	if r.Count("entity") != 0 || r.Errors != 0 || r.byTable == nil {
		t.Fatalf("got %+v", r)
	}
}
