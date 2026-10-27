//ff:func feature=load type=helper control=sequence
//ff:what nulJSON — a string with NUL bytes as its JSON string literal, which text can hold; NULL for any other string
package load

import (
	"encoding/json"
	"strings"
)

// nulJSON returns the JSON string literal of s (NUL written as the escape
// \u0000) when s has a NUL byte, which PostgreSQL text cannot hold, and nil
// otherwise. It goes in the *_json column next to the cleaned text, so the
// exact value is kept.
func nulJSON(s string) any {
	if !strings.ContainsRune(s, 0) {
		return nil
	}
	b, _ := json.Marshal(s)
	return string(b)
}
