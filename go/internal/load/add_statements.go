//ff:func feature=load type=model control=iteration dimension=2
//ff:what Rows.addStatements — the statements of an entity, properties in sorted order, each with its key (statementID)
package load

import (
	"slices"

	"github.com/geul-org/geul-dev/wiki/go/internal/dump"
)

// addStatements adds every statement of e: properties in sorted order (so
// keys are the same on every run), each property's statements in array order
// with that position as the ordinal. The i-th statement gets
// statementID(dump, line, i); past the key range the rest is recorded as one
// load_error.
func (r *Rows) addStatements(c Config, line int64, e dump.Entity, refs *RefSet) {
	props := make([]string, 0, len(e.Claims))
	for p := range e.Claims {
		props = append(props, p)
	}
	slices.Sort(props)
	i := 0
	for _, p := range props {
		for ord, s := range e.Claims[p] {
			sid, err := statementID(c.DumpID, line, i)
			if err != nil {
				r.Error(c.DumpID, line, e.ID, "statement", err.Error(), nil)
				return
			}
			r.addStatement(c, line, e.ID, sid, ord, s, refs)
			i++
		}
	}
}
