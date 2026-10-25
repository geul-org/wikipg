//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestPrecisionName — named precisions and a number for the rest
package port

import "testing"

func TestPrecisionName(t *testing.T) {
	for p, want := range map[int]string{10: "월", 9: "연", 8: "10년", 7: "세기", 6: "천년", 3: "3"} {
		if got := precisionName(p); got != want {
			t.Fatalf("%d: %q", p, got)
		}
	}
}
