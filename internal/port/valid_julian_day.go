//ff:func feature=port type=helper control=sequence
//ff:what validJulianDay — whether month-day exists in a historical year of the proleptic Julian calendar
package port

// validJulianDay reports whether m-d exists in historical year y (no year 0,
// -1 is 1 BCE) of the proleptic Julian calendar, where every fourth
// astronomical year is a leap year: 1 BCE, 5 BCE, …, 4, 8, ….
func validJulianDay(y int64, m, d int) bool {
	if y == 0 || m < 1 || m > 12 || d < 1 {
		return false
	}
	if y < 0 {
		y++
	}
	length := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[m]
	if m == 2 && ((y%4)+4)%4 == 0 {
		length = 29
	}
	return d <= length
}
