//ff:type feature=repair type=model
//ff:what outcome — a worker's answer for one batch: what it found or why it failed
package repair

// outcome is what a worker sends for one batch: the Found or the error that
// stopped it (Found.Seq is set either way).
type outcome struct {
	Found Found
	Err   error
}
