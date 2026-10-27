//ff:type feature=dump type=model
//ff:what Sitelink — a link to a page of a Wikimedia site: site, title, badges
package dump

// Sitelink links an entity to a page: the site (enwiki), the page title and
// the badge items (Q IDs).
type Sitelink struct {
	Site   string   `json:"site"`
	Title  string   `json:"title"`
	Badges []string `json:"badges"`
}
