//ff:func feature=load type=helper control=sequence
//ff:what nullable — a string as a COPY value: NULL when empty
package load

// nullable returns s as a column value, nil (NULL) when it is empty.
func nullable(s string) any {
	if s == "" {
		return nil
	}
	return s
}
