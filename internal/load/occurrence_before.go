//ff:func feature=load type=helper control=sequence
//ff:what Occurrence.Before — whether an occurrence comes before another in dump order (statement key, then reference position)
package load

// Before reports whether o comes before p in dump order: a smaller statement
// key, or the same statement and an earlier reference.
func (o Occurrence) Before(p Occurrence) bool {
	return o.Statement < p.Statement || o.Statement == p.Statement && o.Ordinal < p.Ordinal
}
