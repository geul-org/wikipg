//ff:func feature=port type=helper control=sequence
//ff:what prov — test fixture: a provenance with a short URI and three tools at SHA-1 commits
package port

import "strings"

// prov is a small provenance for tests: loader, porter and library at
// commits 11…, 22… and 33….
func prov() Provenance {
	return Provenance{DumpURI: "d", AcquiredAt: 5, Tools: []Tool{
		{Doc: 3, Commit: strings.Repeat("11", 20)}, {Doc: 1, Commit: strings.Repeat("22", 20)}, {Doc: 2, Commit: strings.Repeat("33", 20)},
	}}
}
