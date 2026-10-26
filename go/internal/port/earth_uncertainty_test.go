//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestEarthUncertainty — matches the Georeferencing Best Practices Table 3 and Example 4; huge precisions are capped
package port

import "testing"

func TestEarthUncertainty(t *testing.T) {
	cases := []struct {
		p, lat float64
		m      uint32
	}{
		{1, 0, 156904}, {1, 60, 124605}, {0.0001, 30, 15}, {1.0 / 3600, 0, 44}, {1.0 / 3600, 85, 32}, {0.01, 10.27, 1557},
	}
	for _, c := range cases {
		if got := (earthUncertainty(c.p, c.lat) + 999) / 1000; got != c.m {
			t.Fatalf("%v° at %v°: %d m want %d", c.p, c.lat, got, c.m)
		}
	}
	if earthUncertainty(1e9, 0) != 0xFFFFFFFE {
		t.Fatal("cap")
	}
}
