//ff:func feature=load type=model control=sequence
//ff:what TestNewRefSet — the bound spreads over shards with at least one per shard
package load

import "testing"

func TestNewRefSet(t *testing.T) {
	if s := NewRefSet(6400); s.limit != 100 || s.shards[63].seen == nil {
		t.Fatalf("limit %d", s.limit)
	}
	if s := NewRefSet(0); s.limit != 1 {
		t.Fatalf("zero: %d", s.limit)
	}
}
