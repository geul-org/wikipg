//ff:func feature=load type=helper control=sequence
//ff:what TestConfigKeepLang — no languages keeps all; chosen languages keep only those
package load

import "testing"

func TestConfigKeepLang(t *testing.T) {
	if !(Config{}).keepLang("de") {
		t.Fatal("all: want true")
	}
	c := Config{Langs: map[string]bool{"ko": true}}
	if !c.keepLang("ko") || c.keepLang("en") {
		t.Fatal("chosen: wrong")
	}
}
