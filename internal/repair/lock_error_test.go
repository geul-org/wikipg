//ff:func feature=repair type=helper control=sequence
//ff:what TestLockError — a query error passes through; no error means another run holds the lock
package repair

import (
	"errors"
	"strings"
	"testing"
)

func TestLockError(t *testing.T) {
	e := errors.New("x")
	if lockError(e) != e {
		t.Fatal("pass through")
	}
	if err := lockError(nil); err == nil || !strings.Contains(err.Error(), "build-reference-order run holds the lock") {
		t.Fatal(err)
	}
}
