//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestEdtfYear — CE years in four digits, BCE shifted by one, large years with Y, year 0 rejected
package port

import "testing"

func TestEdtfYear(t *testing.T) {
	for y, want := range map[int64]string{1952: "1952", 5: "0005", -1: "0000", -44: "-0043", 12000: "Y12000", -13800000000: "Y-13799999999"} {
		if got, ok := edtfYear(y); !ok || got != want {
			t.Fatalf("%d: %q %v want %q", y, got, ok, want)
		}
	}
	if _, ok := edtfYear(0); ok {
		t.Fatal("0: want false")
	}
}
