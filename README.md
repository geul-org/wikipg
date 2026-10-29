# wikipg

`wikipg` loads a Wikidata JSON dump into PostgreSQL without losing anything. Entities, labels, descriptions, aliases, claims, qualifiers, references, sitelinks and Lexemes are stored in dump order. Every dump line is also kept verbatim, and the dump itself (file, URL, size, checksums, loader commit) is recorded in a `dump` row.

wikipg does not depend on GEUL. [wikigeul](https://github.com/geul-org/wikigeul) reads the loaded database and converts it into GEUL documents.

## Install

```sh
go install github.com/geul-org/wikipg/cmd/wikipg@latest
```

## Requirements

- Go 1.25 or later.
- PostgreSQL. The schema is written for PostgreSQL 16 and uses lz4 compression for `entity_raw`. A full dump needs about 2 TB including indexes.
- `lbzip2` (parallel, preferred) or `bzip2` to read `.json.bz2` dumps. Uncompressed `.json` files also work.
- A Wikidata dump from [dumps.wikimedia.org](https://dumps.wikimedia.org/wikidatawiki/entities/) and its SHA-256 next to it as `<file>.sha256` (`sha256sum` format). wikipg refuses a dump without a SHA-256. If Wikimedia's `wikidata-YYYYMMDD-md5sums.txt` and `-sha1sums.txt` are in the same directory, those checksums are recorded too.

## Usage

The database connection comes from `-dsn` or, when it is empty, from the standard libpq environment variables (`PGHOST`, `PGUSER`, `PGDATABASE`, …) and `~/.pgpass`.

1. Create the schema (sections 1–3 of the file; section 4 is commented out):

   ```sh
   psql -f sql/wikidata.sql
   ```

2. Load a dump:

   ```sh
   wikipg -in wikidata-20261019-all.json.bz2 -commit <wikipg commit>
   wikipg -in wikidata-20261021-lexemes.json.bz2 -dump-id 2 -commit <wikipg commit>
   ```

   A load resumes from its checkpoint when run again. Main flags (`wikipg -h` lists all):

   | Flag | Meaning | Default |
   |---|---|---|
   | `-in` | Dump file, `.json` or `.json.bz2` (required) | |
   | `-commit` | Git commit of wikipg, recorded on the dump row (required) | |
   | `-dump-id` | ID of this dump in the database (1–255) | `1` |
   | `-schema` | Target schema | `wikidata` |
   | `-langs` | Languages of labels, descriptions, aliases and glosses: a comma list or `all` | `all` |
   | `-from-line` | First dump line to read (`0`: resume after the checkpoint) | `0` |
   | `-limit` | Load at most this many entity lines (`0`: all) | `0` |
   | `-url`, `-dump-date`, `-published` | Dump metadata; URL and date default from the file name | |
   | `-decoders`, `-parsers`, `-writers`, `-batch-mib` | Parallelism and batch size; defaults leave about 30% of the machine free | |

3. After the load, run section 4 of `sql/wikidata.sql` (uncomment it): resume cleanup, reference deduplication, constraints and indexes.

4. Build the per-occurrence reference snak order:

   ```sh
   wikipg build-reference-order
   ```

`wikipg repair-reference-snak` repairs databases whose `reference_snak` was deduplicated by an older rule. Both subcommands take `-dry-run`, a line range (`-from-line`, `-to-line`) or an entity range (`-entities Q1-Q10000`), and resume from their own checkpoint.

## Documentation

Details (in Korean):

- [docs/schema.md](docs/schema.md) — principles, tables, Lexemes, the load procedure
- [docs/reference-snak.md](docs/reference-snak.md) — how reference snaks are stored, `reference_snak_order`, `build-reference-order` and `repair-reference-snak`

The schema itself, with column comments, is [sql/wikidata.sql](sql/wikidata.sql). Other Go modules can embed it through the package `github.com/geul-org/wikipg/sql`.

## Layout

```
cmd/wikipg/        command: load, build-reference-order, repair-reference-snak
internal/dump/     reading the dump (decompression, lines, entity decoding)
internal/load/     building rows, COPY, dump record and checkpoints
internal/repair/   reference snak repair and per-occurrence order
sql/               schema and the tables of the two subcommands
```

## Development

```sh
go build ./... && go vet ./... && go test ./...
```

Database tests create `sql/wikidata.sql` in a temporary schema through the `PG*` environment variables and are skipped when no database is reachable.

## License

Our own work in this repository is dedicated to the public domain under [CC0 1.0](LICENSE). Third-party components keep their own licenses; see [NOTICE](NOTICE). Wikidata dumps are CC0.
