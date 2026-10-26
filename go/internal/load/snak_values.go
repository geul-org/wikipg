//ff:func feature=load type=helper control=selection
//ff:what Rows.snakValues — the snak columns of a snak: type, datatype, value type, value JSON, entity ID
package load

import (
	"bytes"

	"github.com/geul-org/wikiport/go/internal/dump"
)

// snakValues returns the snak columns (snaktype, datatype, value_type,
// value, value_entity) of s, NULL where absent. The value is the datavalue
// JSON as written. PostgreSQL jsonb rejects the escape \u0000, so such a
// value is NULL and recorded as a load_error of the line (the raw line keeps
// it).
func (r *Rows) snakValues(c Config, line int64, id string, s dump.Snak) []any {
	vals := []any{s.SnakType, nullable(s.Datatype), nil, nil, nil}
	dv := s.DataValue
	switch {
	case dv == nil:
	case bytes.Contains(dv.Value, []byte(`\u0000`)):
		vals[2] = dv.Type
		r.Error(c.DumpID, line, id, "value", `datavalue with \u0000 (jsonb cannot hold it): `+s.Property, nil)
	default:
		vals[2], vals[3], vals[4] = dv.Type, string(dv.Value), nullable(dump.ValueEntity(dv))
	}
	return vals
}
