package domain

import (
	"github.com/ehganzlieb/willfahren/dto"
)

func FilterStops(stops []dto.Stop, maxdistance float64) ImmoListingsFilter {
	return func(il ImmoListing) bool {
		for _, stop := range stops {
			if il.Location == nil {
				return false
			}
			if stop.Location.Distance(*il.Location, dto.DistanceFormulaDefault) <= maxdistance { //TODO: parametrize
				return true
			}
		}
		return false
	}
}
