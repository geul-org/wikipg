//ff:func feature=port type=helper control=selection
//ff:what calendarName — a readable name of a non-Gregorian Wikidata calendar model for the skip report
package port

// calendarName names a calendar model the porter does not convert.
func calendarName(model string) string {
	switch model {
	case julian:
		return "율리우스력"
	default:
		return "달력 " + model
	}
}
