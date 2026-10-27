//ff:type feature=load type=model
//ff:what RefSet — the reference hashes already written, kept within a memory bound, shared by workers
package load

// RefSet remembers reference hashes already written so a source shared by
// many claims is written once. It is split into shards, each with its own
// lock and a bound; a full shard forgets everything and starts over, so
// memory stays bounded and a forgotten hash may be written again (wikidata.sql
// 4.2 removes those duplicates).
type RefSet struct {
	shards [64]refShard
	limit  int
}
