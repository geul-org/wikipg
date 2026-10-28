//ff:func feature=repair type=helper control=sequence
//ff:what lockError — why the repair lock was not taken: the query's error, or another run holding it
package repair

import "errors"

// lockError returns err when the lock query failed, otherwise the error that
// another writing run holds the lock.
func lockError(err error) error {
	if err != nil {
		return err
	}
	return errors.New("repair: another repair-reference-snak run holds the lock")
}
