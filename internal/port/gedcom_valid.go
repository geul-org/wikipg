//ff:func feature=port type=helper control=sequence
//ff:what gedcomValid — whether github.com/cacack/gedcom-go parses and validates a GEDCOM date value as a Julian date
package port

import (
	"errors"

	"github.com/cacack/gedcom-go/v3/gedcom"
)

// gedcomValid checks s with github.com/cacack/gedcom-go/v3 (MIT,
// THIRD_PARTY_NOTICES.md): gedcom.ParseDate, Date.Validate and the Julian
// calendar. Validate checks the day only on Gregorian dates, so on a Julian
// one it checks just the dual year (JULIAN 30 FEB 1100 and JULIAN 40 OCT 1066
// pass); gedcomString checks the day first (validJulianDay).
func gedcomValid(s string) error {
	d, err := gedcom.ParseDate(s)
	if err != nil {
		return err
	}
	if err := d.Validate(); err != nil {
		return err
	}
	if d.Calendar != gedcom.CalendarJulian {
		return errors.New("not a Julian date")
	}
	return nil
}
