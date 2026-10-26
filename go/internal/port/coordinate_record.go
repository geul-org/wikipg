//ff:func feature=port type=encoder control=selection
//ff:what coordinateRecord — a Wikidata globecoordinate as COORDINATE-EARTH or COORDINATE-GEO; otherwise the reason it is not converted
package port

import (
	"encoding/json"
	"math"
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// coordinateRecord converts a globecoordinate datavalue: on Earth (globe Q2)
// to COORDINATE-EARTH with the uncertainty from its precision
// (earthUncertainty), elsewhere to COORDINATE-GEO with the globe as a
// WIKIDATA-QID and the uncertainty unknown (the WGS 84 ellipsoid does not
// fit another body). A missing precision is an unknown uncertainty. A value
// with no or a non-item globe, an altitude (COORDINATE-EARTH holds none),
// degrees out of range or a precision that is not positive and at most 180°
// gives its reason (claimObject keeps such values as
// original text).
func coordinateRecord(value string) (core.Record, string) {
	var v struct {
		Lat       float64  `json:"latitude"`
		Lon       float64  `json:"longitude"`
		Precision *float64 `json:"precision"`
		Altitude  *float64 `json:"altitude"`
		Globe     string   `json:"globe"`
	}
	if err := json.Unmarshal([]byte(value), &v); err != nil {
		return core.Record{}, "globecoordinate: JSON 아님"
	}
	globe, okGlobe := entityRef(strings.TrimPrefix(v.Globe, "http://www.wikidata.org/entity/"))
	earth := v.Globe == "http://www.wikidata.org/entity/Q2"
	lonMax := 360.0
	if earth {
		lonMax = 180
	}
	c := core.Coordinate{Lat: v.Lat, Lon: v.Lon, U: core.UncertaintyUnknown}
	switch {
	case !okGlobe || globe.Codec != core.RefWikidataQID<<3:
		return core.Record{}, "globecoordinate: 천체 없음"
	case v.Altitude != nil:
		return core.Record{}, "globecoordinate: 고도"
	case v.Lat < -90 || v.Lat > 90 || v.Lon < -180 || v.Lon > lonMax:
		return core.Record{}, "globecoordinate: 범위"
	case v.Precision != nil && (!(*v.Precision > 0) || *v.Precision > 180 || math.IsInf(*v.Precision, 0)):
		return core.Record{}, "globecoordinate: 정밀도"
	case earth && v.Precision != nil:
		c.U = earthUncertainty(*v.Precision, v.Lat)
	case !earth:
		c.Globe = &globe
	}
	return core.EncodeCoordinate(c), ""
}
