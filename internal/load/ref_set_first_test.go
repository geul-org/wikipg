//ff:func feature=load type=helper control=iteration dimension=1
//ff:what TestRefSetFirst — a hash is first once; a full shard forgets and a hash may be first again
package load

import (
	"fmt"
	"testing"
)

func TestRefSetFirst(t *testing.T) {
	s := NewRefSet(6400)
	if !s.First("r1") || s.First("r1") {
		t.Fatal("first once")
	}
	small := NewRefSet(1)
	for i := 0; i < 200; i++ {
		small.First(fmt.Sprint("h", i))
	}
	if !small.First("h0") {
		t.Fatal("forgotten hash should be first again")
	}
}
