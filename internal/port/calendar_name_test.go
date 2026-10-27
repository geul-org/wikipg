//ff:func feature=port type=helper control=sequence
//ff:what TestCalendarName — an unconverted model is named by its URI
package port

import "testing"

func TestCalendarName(t *testing.T) {
	if calendarName("x") != "달력 x" {
		t.Fatal("wrong")
	}
}
