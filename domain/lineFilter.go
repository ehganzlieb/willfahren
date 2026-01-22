package domain

import (
	"github.com/ehganzlieb/2025-10-26_willfahren/dto"
)

func FilterStops(stops []dto.Stop, maxdistance float64) ImmoListingsFilter {
	return func(il ImmoListing) bool {
		for _, stop := range stops {
			if stop.Location.Distance(il.Location) <= maxdistance {
				return true
			}
		}
		return false
	}
}
