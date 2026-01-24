package domain

import (
	"slices"
	"strings"

	"github.com/ehganzlieb/2025-10-26_willfahren/dto"
)

type ImmoListing dto.Apartment
type ImmoListings []ImmoListing

type ImmoListingsFilter func(ImmoListing) bool

func InvertImmoListingsFilter(filter ImmoListingsFilter) ImmoListingsFilter {
	return func(il ImmoListing) bool {
		return !filter(il)
	}
}

func MergeFilters(filters ...ImmoListingsFilter) ImmoListingsFilter {
	return func(il ImmoListing) bool {
		for _, f := range filters {
			if !f(il) {
				return false
			}
		}
		return true
	}
}

func (il ImmoListings) ApplyFilter(filter ImmoListingsFilter) ImmoListings {
	var ls []ImmoListing = []ImmoListing(il)
	var lc []ImmoListing = slices.Clone(ls)
	invFilter := InvertImmoListingsFilter(filter)
	return slices.DeleteFunc(lc, invFilter)
}

func FilterDistricts(districts []dto.District) ImmoListingsFilter {
	return func(il ImmoListing) bool {
		return slices.ContainsFunc(districts, func(d dto.District) bool {
			return il.District.PostCode() == d.PostCode()
		})
	}
}

func FilterRooms(minRooms, maxRooms int) ImmoListingsFilter {
	if minRooms == 0 && maxRooms == 0 {
		return func(il ImmoListing) bool {
			return true
		}
	}
	if minRooms == 0 {
		return func(il ImmoListing) bool {
			return il.Rooms <= float32(maxRooms)
		}
	}
	if maxRooms == 0 {
		return func(il ImmoListing) bool {
			return il.Rooms >= float32(minRooms)
		}
	}
	return func(il ImmoListing) bool {
		return il.Rooms >= float32(minRooms) && il.Rooms <= float32(maxRooms)
	}
}

func FilterPrice(minPrice, maxPrice float32) ImmoListingsFilter {
	if minPrice == 0 && maxPrice == 0 {
		return func(il ImmoListing) bool {
			return true
		}
	}
	if minPrice == 0 {
		return func(il ImmoListing) bool {
			return il.Price <= maxPrice
		}
	}
	if maxPrice == 0 {
		return func(il ImmoListing) bool {
			return il.Price >= minPrice
		}
	}
	return func(il ImmoListing) bool {
		return il.Price >= minPrice && il.Price <= maxPrice
	}
}

func FilterArea(minArea, maxArea float32) ImmoListingsFilter {
	if minArea == 0 && maxArea == 0 {
		return func(il ImmoListing) bool {
			return true
		}
	}
	if minArea == 0 {
		return func(il ImmoListing) bool {
			return il.Area <= maxArea
		}
	}
	if maxArea == 0 {
		return func(il ImmoListing) bool {
			return il.Area >= minArea
		}
	}
	return func(il ImmoListing) bool {
		return il.Area >= minArea && il.Area <= maxArea
	}
}

func FilterKeyWords(keywords []string) ImmoListingsFilter {
	for i := range keywords {
		keywords[i] = strings.ToLower(keywords[i])
	}
	return func(il ImmoListing) bool {
		for _, kw := range keywords {
			if strings.Contains(strings.ToLower(il.Description), kw) {
				return true
			}
			if strings.Contains(strings.ToLower(il.Description), kw) {
				return true
			}
		}
		return false
	}
}
