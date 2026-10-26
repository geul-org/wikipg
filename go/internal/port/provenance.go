//ff:type feature=port type=model
//ff:what Provenance — the source and the tools every ported document records in its HEAD
package port

// Provenance is what every ported document records in its HEAD (wiki/README
// 출처 기록): the dump URI (SOURCE), when the dump file was acquired
// (ACQUIRED-AT, POSIX seconds) and the tools that turned it into GEUL, one
// TRANSFORMED-BY each in order — the PostgreSQL loader, this porter and the
// GEUL library it builds with.
type Provenance struct {
	DumpURI    string
	AcquiredAt int64
	Tools      []Tool
}
