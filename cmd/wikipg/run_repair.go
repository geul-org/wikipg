//ff:func feature=cli type=command control=sequence
//ff:what runRepair — wikipg repair-reference-snak: repairs (or with -dry-run compares) reference_snak against entity_raw and reports; stops cleanly on interrupt
package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/lib/pq"

	"github.com/geul-org/wikipg/internal/repair"
)

// runRepair parses args (parseRepairFlags), opens the database (sessions
// with jit off, as the per-line and per-hash lookups are cheap but estimated
// large; sequential scans off, as every query of the repair goes by an index
// and a scan of these tables takes hours; no parallel workers, so the
// queries at once are at most the workers; read-only for -dry-run: dsnWith;
// a DSN lib/pq cannot parse fails here) with a connection per worker and two
// more, runs repair.Run until done or SIGINT/SIGTERM (batches already read
// are still applied, so the checkpoint stays exact) and writes the report
// to stdout.
func runRepair(args []string) error {
	f, err := parseRepairFlags(args)
	if err != nil {
		return err
	}
	dsn := dsnWith(dsnWith(dsnWith(f.DSN, "jit=off"), "enable_seqscan=off"), "max_parallel_workers_per_gather=0")
	if f.Config.DryRun {
		dsn = dsnWith(dsn, "default_transaction_read_only=on")
	}
	conn, err := pq.NewConnector(dsn)
	if err != nil {
		return err
	}
	db := sql.OpenDB(conn)
	defer db.Close()
	db.SetMaxOpenConns(f.Config.Workers + 2)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	st, err := repair.Run(ctx, db, f.Config)
	writeReport(os.Stdout, f.Config, st)
	return err
}
