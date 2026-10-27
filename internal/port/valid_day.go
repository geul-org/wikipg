//ff:func feature=port type=helper control=sequence
//ff:what validDay — whether month and day name a real day of the proleptic Gregorian year
package port

// validDay reports whether m-d exists in year y of the proleptic Gregorian
// calendar (1 ≤ m ≤ 12, 1 ≤ d ≤ the month's length, 29 February only in leap years).
func validDay(y int64, m, d int) bool {
	if m < 1 || m > 12 || d < 1 {
		return false
	}
	length := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[m]
	leap := y%4 == 0 && (y%100 != 0 || y%400 == 0)
	if m == 2 && leap {
		length = 29
	}
	return d <= length
}
