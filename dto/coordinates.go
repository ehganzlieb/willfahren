package dto

import (
	"math"

	"github.com/asmarques/geodist"
)

type Coordinates struct {
	X, Y float64 //in degrees
}

func (c *Coordinates) toGeoDistPoint() geodist.Point {
	return geodist.Point{Long: c.X, Lat: c.Y}
}

type DistanceFormula int

const (
	DistanceFormulaManhattan DistanceFormula = iota //simplest, good approximation of walking distance in cities
	DistanceFormulaHaversine
	DistanceFormulaVincenty //high computational cost, if an error occurs, it falls back to DistanceFormulaHaversine

	DistanceFormulaDefault = DistanceFormulaHaversine
)

/*
Distance() calculates the distance between two coordinates using a given distance formula.
If the formula is DistanceFormulaVincenty and an error occurs, it falls back to DistanceFormulaHaversine.
If no valid formula is given, it defaults to DistanceFormulaDefault.
*/
func (c Coordinates) Distance(other Coordinates, formula DistanceFormula) float64 {
	switch formula {
	case DistanceFormulaManhattan:
		return c.ManhattanDistance(other)
	case DistanceFormulaHaversine:
		return c.HaversineDistance(other)
	case DistanceFormulaVincenty:
		dist, err := c.VincentyDistance(other)
		if err != nil {
			return c.HaversineDistance(other)
		} else {
			return dist
		}
	default:
		return c.Distance(other, DistanceFormulaDefault)
	}
}

// calculate the difference in longitude and latitude in meters using the Manhattan formula
func (c Coordinates) ManhattanDistance(other Coordinates) float64 {
	latitudeDifference := math.Abs(c.Y-other.Y) * 111.320
	longitudeDifference := math.Abs(c.X-other.X) * 111.320 * math.Cos(c.Y*math.Pi/180)

	return latitudeDifference + longitudeDifference

}

func (c Coordinates) HaversineDistance(other Coordinates) float64 {
	return geodist.HaversineDistance(c.toGeoDistPoint(), other.toGeoDistPoint())
}

func (c Coordinates) VincentyDistance(other Coordinates) (float64, error) {
	return geodist.VincentyDistance(c.toGeoDistPoint(), other.toGeoDistPoint())
}
