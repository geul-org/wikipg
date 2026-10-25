//ff:type feature=port type=model
//ff:what wikidataTime — a Wikidata time datavalue as the loader stored it (JSON)
package port

// wikidataTime is a Wikidata time datavalue: an ISO-like signed timestamp,
// its precision (11 = day, 10 = month, 9 = year, …), the timezone offset in
// minutes, the uncertainty bounds and the calendar model entity URI.
type wikidataTime struct {
	Time          string `json:"time"`
	Timezone      int    `json:"timezone"`
	Before        int    `json:"before"`
	After         int    `json:"after"`
	Precision     int    `json:"precision"`
	CalendarModel string `json:"calendarmodel"`
}
