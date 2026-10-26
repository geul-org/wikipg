//ff:func feature=load type=helper control=sequence
//ff:what Config.keepLang — whether a term language is kept (every language when none is chosen)
package load

// keepLang reports whether terms in lang are written: every language when
// Langs is empty, else only the chosen ones.
func (c Config) keepLang(lang string) bool {
	return len(c.Langs) == 0 || c.Langs[lang]
}
