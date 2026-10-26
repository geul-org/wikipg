//ff:func feature=cli type=helper control=sequence
//ff:what fixtureDump — test fixture: the dump row of testdata/fixture.sql as provenanceOf reads it
package main

import "github.com/geul-org/wikiport/go/internal/source"

// fixtureDump is a dump loaded by wikiload, for building a provenance.
func fixtureDump() source.Dump {
	return source.Dump{URL: "u", Loader: wikiloadPath, LoaderCommit: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
}
