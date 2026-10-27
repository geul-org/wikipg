//ff:func feature=port type=helper control=selection
//ff:what rankWord — the TRIPLE rank code of a Wikidata rank: normal 0, preferred 1, deprecated 2; false for others
package port

import "github.com/geul-org/gogeul/pkg/core"

// rankWord maps a Wikidata rank to the TRIPLE rank code (edge/TRIPLE ENUM):
// normal 0, preferred 1, deprecated 2.
func rankWord(rank string) (core.Word, bool) {
	switch rank {
	case "normal":
		return 0, true
	case "preferred":
		return 1, true
	case "deprecated":
		return 2, true
	default:
		return 0, false
	}
}
