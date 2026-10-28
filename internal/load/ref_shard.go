//ff:type feature=load type=model
//ff:what refShard — one shard of a RefSet: its lock and the hashes it remembers with their earliest occurrence written
package load

import "sync"

// refShard is one part of a RefSet, holding the hashes that fall into it,
// each with the earliest occurrence written so far, under its own lock.
type refShard struct {
	mu   sync.Mutex
	seen map[string]Occurrence
}
