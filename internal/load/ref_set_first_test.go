//ff:func feature=load type=helper control=iteration dimension=1
//ff:what TestRefSetFirst — a hash is written at its first occurrence seen and again only at an earlier one; a full shard forgets; a nil set writes all
package load

import (
	"fmt"
	"testing"
)

func TestRefSetFirst(t *testing.T) {
	s := NewRefSet(6400)
	at := func(sid int64, ord int) Occurrence { return Occurrence{Statement: sid, Ordinal: ord} }
	if !s.First("r1", at(20, 1)) || s.First("r1", at(20, 1)) || s.First("r1", at(30, 0)) {
		t.Fatal("same or later occurrence must not be written")
	}
	if !s.First("r1", at(20, 0)) || !s.First("r1", at(10, 5)) || s.First("r1", at(20, 0)) {
		t.Fatal("an earlier occurrence must be written, then remembered")
	}
	small := NewRefSet(1)
	for i := 0; i < 200; i++ {
		small.First(fmt.Sprint("h", i), at(int64(i), 0))
	}
	if !small.First("h0", at(500, 0)) {
		t.Fatal("forgotten hash should be first again")
	}
	var none *RefSet
	if !none.First("r1", at(1, 0)) || !none.First("r1", at(1, 0)) {
		t.Fatal("nil set writes every occurrence")
	}
}
