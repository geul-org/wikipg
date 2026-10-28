//ff:type feature=load type=model
//ff:what RefSet — the reference hashes already written with their earliest occurrence, kept within a memory bound, shared by workers
package load

// RefSet remembers reference hashes already written, each with the earliest
// occurrence written, so a source shared by many claims is written once or a
// few times rather than at every occurrence. It is split into shards, each
// with its own lock and a bound; a full shard forgets everything and starts
// over, so memory stays bounded and a forgotten hash may be written again.
// Parallel workers see occurrences out of dump order, so an occurrence
// earlier than the remembered one is written too: the first occurrence of
// every hash is always written, and wikidata.sql 4.2 keeps exactly that one.
// A nil RefSet writes every occurrence.
type RefSet struct {
	shards [64]refShard
	limit  int
}
