//ff:func feature=repair type=helper control=selection
//ff:what Config.Run — the name of the run's range, with the dump the key of its checkpoint (lines:2-0, entities:Q1-Q10000)
package repair

import "fmt"

// Run names the range of c, the key its checkpoint is kept under with c.DumpID: lines:From-To
// (To 0: to the end) or entities:KindLo-KindHi.
func (c Config) Run() string {
	switch c.Kind {
	case "":
		return fmt.Sprintf("lines:%d-%d", c.From, c.To)
	default:
		return fmt.Sprintf("entities:%s%d-%s%d", c.Kind, c.Lo, c.Kind, c.Hi)
	}
}
