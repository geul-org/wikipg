//ff:func feature=repair type=helper control=sequence
//ff:what TestNullText — a string is valid, nil is NULL
package repair

import (
	"database/sql"
	"testing"
)

func TestNullText(t *testing.T) {
	if nullText("a") != (sql.NullString{String: "a", Valid: true}) || nullText(nil).Valid {
		t.Fatal("null text")
	}
}
