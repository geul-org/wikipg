//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what Stats.addOrder — adds one applied batch of an order build: its counts, the occurrence orders found and written, examples
package repair

import "fmt"

// addOrder adds the applied batch f of an order build to st: the counts of
// add (no hash is processed), the occurrences in another order (Orders,
// Written when stored) and up to 20 examples of them.
func (st *Stats) addOrder(f Found, written bool) {
	st.add(f, nil, 0, false)
	st.Orders += int64(len(f.Orders))
	if written {
		st.Written += int64(len(f.Orders))
	}
	for _, o := range f.Orders[:min(len(f.Orders), max(20-len(st.Examples), 0))] {
		st.Examples = append(st.Examples, fmt.Sprintf("%s %s order %v (statement %d reference %d)", o.Hash, o.Entity, o.Perm, o.At.Statement, o.At.Ordinal))
	}
}
