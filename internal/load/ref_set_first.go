//ff:func feature=load type=helper control=sequence
//ff:what RefSet.First — whether a reference occurrence must be written: its hash is new or this occurrence is earlier than the one remembered (and remembers it)
package load

import "hash/fnv"

// First reports whether the occurrence at of hash must be written: the hash
// has not been seen (within the bound), or at comes before the earliest
// occurrence remembered for it. It then remembers at. A shard that reaches
// its bound is cleared first. A nil RefSet answers true for every occurrence.
func (s *RefSet) First(hash string, at Occurrence) bool {
	if s == nil {
		return true
	}
	f := fnv.New32a()
	f.Write([]byte(hash))
	sh := &s.shards[f.Sum32()%uint32(len(s.shards))]
	sh.mu.Lock()
	defer sh.mu.Unlock()
	prev, ok := sh.seen[hash]
	if ok && !at.Before(prev) {
		return false
	}
	if !ok && len(sh.seen) >= s.limit {
		sh.seen = map[string]Occurrence{}
	}
	sh.seen[hash] = at
	return true
}
