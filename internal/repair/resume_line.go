//ff:func feature=repair type=reader control=sequence
//ff:what resumeLine — the line a run starts after: its own dump's checkpoint when writing, 0 for a dry run (which still checks the key and reports the checkpoint) or without a checkpoint table
package repair

import (
	"database/sql"
	"fmt"
)

// resumeLine checks the key of the checkpoint table of c (checkKey) and
// returns the line the run starts after: 0 when the table does not exist (a
// dry run before any writing run), otherwise the checkpoint of run c.Run() of
// dump c.DumpID, written to c.Log. A dry run reports it but starts at the
// range's beginning (0).
func resumeLine(db *sql.DB, c Config) (int64, error) {
	exists, err := checkKey(db, c)
	if err != nil || !exists {
		return 0, err
	}
	line, err := readCheckpoint(db, c)
	if err != nil {
		return 0, err
	}
	if c.DryRun {
		fmt.Fprintf(c.Log, "checkpoint %s dump %d: line %d (dry run: starts at the range's beginning)\n", c.Run(), c.DumpID, line)
		return 0, nil
	}
	fmt.Fprintf(c.Log, "checkpoint %s dump %d: line %d (resuming after it)\n", c.Run(), c.DumpID, line)
	return line, nil
}
