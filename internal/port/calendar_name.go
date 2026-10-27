//ff:func feature=port type=helper control=sequence
//ff:what calendarName — a readable name of an unconverted Wikidata calendar model for the skip report
package port

// calendarName names a calendar model the porter does not convert.
func calendarName(model string) string {
	return "달력 " + model
}
