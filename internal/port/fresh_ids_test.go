//ff:func feature=port type=helper control=sequence
//ff:what freshIDs — test fixture: a LOCAL ID assignment with no previous version, starting at next
package port

import "github.com/geul-org/gogeul/pkg/delta"

// freshIDs is delta.NewLocalIDs for a first version whose IDs start at next.
func freshIDs(next uint64) *delta.LocalIDs {
	ids, _ := delta.NewLocalIDs(nil, next)
	return ids
}
