//ff:func feature=load type=helper control=sequence
//ff:what RefSet.First — whether a reference hash is new (and remembers it)
package load

import "hash/fnv"

// First reports whether hash has not been seen (within the bound) and
// remembers it. A shard that reaches its bound is cleared first.
func (s *RefSet) First(hash string) bool {
	f := fnv.New32a()
	f.Write([]byte(hash))
	sh := &s.shards[f.Sum32()%uint32(len(s.shards))]
	sh.mu.Lock()
	defer sh.mu.Unlock()
	if _, ok := sh.seen[hash]; ok {
		return false
	}
	if len(sh.seen) >= s.limit {
		sh.seen = map[string]struct{}{}
	}
	sh.seen[hash] = struct{}{}
	return true
}
