//ff:func feature=cli type=command control=iteration dimension=1
//ff:what parseFlags — the run config from command-line arguments; full porter and library commits are required
package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/geul-org/wikiport/go/internal/port"
)

// parseFlags reads args into a config. -commit and -lib-commit (the full git
// object IDs of the porter and the GEUL library, written in every document's
// TRANSFORMED-BY) are required and must be well formed (port.CommitDigest).
func parseFlags(args []string) (config, error) {
	fs := flag.NewFlagSet("wikiport", flag.ContinueOnError)
	var c config
	var langs string
	fs.StringVar(&c.Root, "root", ".", "t002 root (holds .well-known/geul)")
	fs.Uint64Var(&c.From, "from", 1, "first QID number")
	fs.Uint64Var(&c.To, "to", 1000, "last QID number")
	fs.StringVar(&langs, "langs", "ko,en", "label and alias languages")
	fs.StringVar(&c.DSN, "dsn", "", "PostgreSQL DSN (empty: PG* environment variables)")
	fs.StringVar(&c.Commit, "commit", "", "porter git commit, full ID (required)")
	fs.StringVar(&c.LibCommit, "lib-commit", "", "GEUL library git commit, full ID (required)")
	fs.StringVar(&c.Report, "report", "", "report TSV path (empty: none)")
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	for _, id := range []string{c.Commit, c.LibCommit} {
		if _, err := port.CommitDigest(id); err != nil {
			return config{}, fmt.Errorf("-commit and -lib-commit: %w", err)
		}
	}
	if c.From == 0 || c.From > c.To {
		return config{}, errors.New("wikiport: need 1 <= -from <= -to")
	}
	c.Langs = strings.Split(strings.ReplaceAll(langs, " ", ""), ",")
	return c, nil
}
