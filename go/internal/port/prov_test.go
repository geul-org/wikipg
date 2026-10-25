//ff:func feature=port type=helper control=sequence
//ff:what prov — test fixture: a provenance with short URIs and versions
package port

// prov is a small provenance for tests.
func prov() Provenance {
	return Provenance{DumpURI: "d", AcquiredAt: 5, Loader: Tool{URI: "l", Version: "1"}, Porter: Tool{URI: "p", Version: "2"}}
}
