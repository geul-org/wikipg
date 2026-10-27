//ff:func feature=cli type=command control=sequence
//ff:what main — wikiload entry: loads a Wikidata JSON dump into the wikidata schema of geuldev (sql/wikidata.sql)
package main

import (
	"fmt"
	"os"
)

// main runs wikiload with the command-line arguments and exits 1 on error.
func main() {
	if err := runLoad(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "wikiload:", err)
		os.Exit(1)
	}
}
