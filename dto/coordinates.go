package dto

import (
	"math"

	"github.com/asmarques/geodist"
)

const (
	EarthMagicNumber = 111.320
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

/*
ManhattanDistance() calculates the Manhattan distance between two coordinates.

The Manhattan distance is the sum of the absolute differences of their respective coordinates.
The result is in meters.
Note that the Manhattan distance is a simple approximation of the walking distance in cities and does not take into account the curvature of the Earth.
For a more accurate approximation, use HaversineDistance() or VincentyDistance().
*/
func (c Coordinates) ManhattanDistance(other Coordinates) float64 {
	latitudeDifference := math.Abs(c.Y-other.Y) * EarthMagicNumber
	longitudeDifference := math.Abs(c.X-other.X) * EarthMagicNumber * math.Cos(c.Y*math.Pi/180)

	return latitudeDifference + longitudeDifference

}

/*
HaversineDistance() calculates the distance between two coordinates using the Haversine formula.

The Haversine formula is an approximation to the great-circle distance between two points on a sphere.
The result is in meters.
*/
func (c Coordinates) HaversineDistance(other Coordinates) float64 {
	return geodist.HaversineDistance(c.toGeoDistPoint(), other.toGeoDistPoint())
}

func (c Coordinates) VincentyDistance(other Coordinates) (float64, error) {
	return geodist.VincentyDistance(c.toGeoDistPoint(), other.toGeoDistPoint())
}
