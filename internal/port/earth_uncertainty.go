//ff:func feature=port type=helper control=sequence
//ff:what earthUncertainty — the uncertainty (mm) a coordinate precision in degrees gives at a latitude on WGS 84 (Georeferencing Best Practices 3.4.3)
package port

import "math"

// earthUncertainty is the uncertainty radius in millimetres that coordinate
// precision p (degrees) gives at latitude lat on the WGS 84 ellipsoid, after
// Chapman & Wieczorek (2020), Georeferencing Best Practices 3.4.3: the
// latitude error p·M(φ)·π/180 and the longitude error p·N(φ)·cos φ·π/180
// combined as √(e_φ² + e_λ²), rounded up. M and N are the meridional and
// prime vertical radii of curvature. The result is capped below the "unknown"
// value. This is the part of the uncertainty that comes from the precision
// alone, a lower bound of the whole (value/numeric/COORDINATE-EARTH).
func earthUncertainty(p, lat float64) uint32 {
	const a, f = 6378137.0, 1 / 298.257223563
	e2 := f * (2 - f)
	s := math.Sin(lat * math.Pi / 180)
	w := 1 - e2*s*s
	m := a * (1 - e2) / math.Pow(w, 1.5)
	n := a / math.Sqrt(w)
	eLat := p * m * math.Pi / 180
	eLon := p * n * math.Cos(lat*math.Pi/180) * math.Pi / 180
	mm := math.Ceil(math.Hypot(eLat, eLon) * 1000)
	return uint32(math.Min(mm, 0xFFFFFFFE))
}
