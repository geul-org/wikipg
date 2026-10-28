//ff:func feature=repair type=helper control=sequence
//ff:what nullText — a loader column value (string or nil) as a NullString
package repair

import "database/sql"

// nullText returns v, a string or nil (NULL) as the loader gives column
// values, as a NullString.
func nullText(v any) sql.NullString {
	s, ok := v.(string)
	return sql.NullString{String: s, Valid: ok}
}
