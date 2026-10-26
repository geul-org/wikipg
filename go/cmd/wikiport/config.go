//ff:type feature=cli type=model
//ff:what config — the wikiport run settings: tree root, QID range, languages, DB DSN, porter and library commits, report path
package main

// config is one wikiport run: the t002 root whose .well-known/geul receives
// the documents, the QID range [From, To], the label/alias languages, the
// PostgreSQL DSN (empty: lib/pq reads the PG* environment variables), the
// porter and GEUL library commits recorded in TRANSFORMED-BY and where the
// report goes.
type config struct {
	Root      string
	From      uint64
	To        uint64
	Langs     []string
	DSN       string
	Commit    string
	LibCommit string
	Report    string
}
