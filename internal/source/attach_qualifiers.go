//ff:func feature=source type=loader control=iteration dimension=1
//ff:what attachQualifiers — puts each qualifier on the claim with its statement ID, keeping their order
package source

// attachQualifiers appends each qualifier in qs to the claim in claims whose
// ID is its statement ID, in qs order. A qualifier of no listed claim is dropped.
func attachQualifiers(claims []Claim, qs []Qualifier) {
	at := map[int64]int{}
	for i, c := range claims {
		at[c.ID] = i
	}
	for _, q := range qs {
		if i, ok := at[q.StatementID]; ok {
			claims[i].Qualifiers = append(claims[i].Qualifiers, q)
		}
	}
}
