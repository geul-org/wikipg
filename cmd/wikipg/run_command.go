//ff:func feature=cli type=command control=selection
//ff:what runCommand — dispatches wikipg's arguments: the repair-reference-snak subcommand, or a load (no subcommand)
package main

// runCommand runs the subcommand args[0] names with the rest of args
// (repair-reference-snak: runRepair); without a subcommand args are the load
// flags (runLoad).
func runCommand(args []string) error {
	switch {
	case len(args) > 0 && args[0] == "repair-reference-snak":
		return runRepair(args[1:])
	default:
		return runLoad(args)
	}
}
