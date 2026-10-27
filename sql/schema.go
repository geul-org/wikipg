// Package sql gives the wikidata schema (wikidata.sql) to other modules, such
// as gogeulwiki's tests, without copying it.
package sql

import _ "embed"

// Schema is wikidata.sql: sections 1–3 create the schema wikidata and its
// tables; section 4 (cleanup, keys and indexes after a load) is commented out.
//
//go:embed wikidata.sql
var Schema string
