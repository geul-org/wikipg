//ff:func feature=cli type=command control=sequence
//ff:what parseFlags — the run config from command-line arguments; the porter commit is required
package main

import (
	"errors"
	"flag"
	"strings"
)

// parseFlags reads args into a config. -commit (the porter's git commit,
// written in every document's TRANSFORMED-BY) is required.
func parseFlags(args []string) (config, error) {
	fs := flag.NewFlagSet("wikiport", flag.ContinueOnError)
	var c config
	var langs string
	fs.StringVar(&c.Root, "root", ".", "t002 root (holds .well-known/geul)")
	fs.Uint64Var(&c.From, "from", 1, "first QID number")
	fs.Uint64Var(&c.To, "to", 1000, "last QID number")
	fs.StringVar(&langs, "langs", "ko,en", "label and alias languages")
	fs.StringVar(&c.DSN, "dsn", "", "PostgreSQL DSN (empty: PG* environment variables)")
	fs.StringVar(&c.Commit, "commit", "", "porter git commit (required)")
	fs.StringVar(&c.Report, "report", "", "report TSV path (empty: none)")
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	if c.Commit == "" {
		return config{}, errors.New("wikiport: -commit is required")
	}
	if c.From == 0 || c.From > c.To {
		return config{}, errors.New("wikiport: need 1 <= -from <= -to")
	}
	c.Langs = strings.Split(langs, ",")
	return c, nil
}
