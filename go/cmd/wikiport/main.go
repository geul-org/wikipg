//ff:func feature=cli type=command control=sequence
//ff:what main — wikiport entrypoint: run, print the error and exit 1 on failure
package main

import (
	"fmt"
	"os"
)

// main runs wikiport (wiki/README) and exits 1 on an error.
func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
