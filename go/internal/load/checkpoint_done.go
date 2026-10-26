//ff:func feature=load type=model control=iteration dimension=1
//ff:what Checkpoint.Done — marks a batch done and advances the line while batches are done in order
package load

// Done marks batch seq, whose last line is last, as done and advances Line
// over every batch done in order from the next expected one. It reports
// whether Line moved.
func (c *Checkpoint) Done(seq, last int64) bool {
	c.done[seq] = last
	moved := false
	for {
		l, ok := c.done[c.next]
		if !ok {
			return moved
		}
		delete(c.done, c.next)
		c.Line, c.next, moved = l, c.next+1, true
	}
}
