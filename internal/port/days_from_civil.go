//ff:func feature=port type=helper control=sequence
//ff:what daysFromCivil — the proleptic Gregorian day count from 1970-01-01 of a year, month and day
package port

// daysFromCivil returns the number of days from 1970-01-01 to y-m-d in the
// proleptic Gregorian calendar (H. Hinnant's days_from_civil). Wikidata year
// numbers are astronomical when the sign is +; Wikidata writes 1 BCE as -0001,
// which this treats as year -1 (the caller rejects BCE dates).
func daysFromCivil(y int64, m, d int) int64 {
	if m <= 2 {
		y--
	}
	era := y / 400
	if y < 0 && y%400 != 0 {
		era--
	}
	yoe := y - era*400
	mp := (int64(m) + 9) % 12
	doy := (153*mp+2)/5 + int64(d) - 1
	doe := yoe*365 + yoe/4 - yoe/100 + doy
	return era*146097 + doe - 719468
}
