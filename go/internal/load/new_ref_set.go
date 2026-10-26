//ff:func feature=load type=model control=iteration dimension=1
//ff:what NewRefSet — an empty RefSet that keeps about limit hashes in all
package load

// NewRefSet returns an empty RefSet keeping about limit hashes in all
// (limit/64 per shard, at least 1).
func NewRefSet(limit int) *RefSet {
	s := &RefSet{limit: max(limit/64, 1)}
	for i := range s.shards {
		s.shards[i].seen = map[string]struct{}{}
	}
	return s
}
