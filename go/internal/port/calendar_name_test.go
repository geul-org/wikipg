//ff:func feature=port type=helper control=sequence
//ff:what TestCalendarName — the Julian model by name, any other model by URI
package port

import "testing"

func TestCalendarName(t *testing.T) {
	if calendarName(julian) != "율리우스력" || calendarName("x") != "달력 x" {
		t.Fatal("wrong")
	}
}
