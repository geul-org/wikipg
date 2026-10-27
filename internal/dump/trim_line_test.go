//ff:func feature=dump type=decoder control=iteration dimension=1
//ff:what TestTrimLine — entity lines lose their comma and newline; brackets and blank lines are false
package dump

import "testing"

func TestTrimLine(t *testing.T) {
	for in, want := range map[string]string{"{\"id\":\"Q1\"},\n": `{"id":"Q1"}`, "{\"id\":\"Q2\"}\r\n": `{"id":"Q2"}`, "{\"id\":\"Q3\"}": `{"id":"Q3"}`} {
		if got, ok := TrimLine([]byte(in)); !ok || string(got) != want {
			t.Fatalf("%q: %q %v", in, got, ok)
		}
	}
	for _, in := range []string{"[\n", "]\n", "\n", ""} {
		if _, ok := TrimLine([]byte(in)); ok {
			t.Fatalf("%q: want false", in)
		}
	}
}
