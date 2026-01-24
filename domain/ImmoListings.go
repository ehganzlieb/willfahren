package domain

import (
	"slices"
	"strings"

	"github.com/ehganzlieb/willfahren/dto"
)

type ImmoListing dto.Apartment
type ImmoListings []ImmoListing

type ImmoListingsFilter func(ImmoListing) bool // a function that takes an ImmoListing and returns true if the ImmoListing passes the filter

/*
InvertImmoListingsFilter inverts the given ImmoListingsFilter.
It takes a filter function that takes an ImmoListing and returns true if the ImmoListing passes the filter and false otherwise.
The returned filter function takes an ImmoListing and returns true if the ImmoListing does not pass the filter and false otherwise.
*/
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

/*
ApplyFilter applies the given filter to the ImmoListings.
The filter is applied by deleting all elements from the ImmoListings
that do not match the filter. The original ImmoListings is not modified.
The modified ImmoListings is returned.
*/
func (il ImmoListings) ApplyFilter(filter ImmoListingsFilter) ImmoListings {
	var ls []ImmoListing = []ImmoListing(il)
	var lc []ImmoListing = slices.Clone(ls)
	invFilter := InvertImmoListingsFilter(filter)
	return slices.DeleteFunc(lc, invFilter)
}

/*
FilterDistricts returns a filter function that filters ImmoListings
based on their district. The filter function takes a slice of
dto.District objects and returns true if the ImmoListing's district
is in the slice, false otherwise.
*/
func FilterDistricts(districts []dto.District) ImmoListingsFilter {
	return func(il ImmoListing) bool {
		return slices.ContainsFunc(districts, func(d dto.District) bool {
			return il.District.PostCode() == d.PostCode()
		})
	}
}

/*
FilterRooms returns a filter function that filters ImmoListings
based on their number of rooms. The filter function takes two
int arguments, minRooms and maxRooms, and returns true if the
ImmoListing's number of rooms is within the range of
[minRooms, maxRooms]. If minRooms is 0, it filters listings
with a number of rooms less than or equal to maxRooms. If
maxRooms is 0, it filters listings with a number of rooms
greater than or equal to minRooms. If both minRooms and
maxRooms are 0, it returns a filter function that always
returns true. If minRooms and maxRooms are both non-zero, it
filters listings with a number of rooms greater than or equal
to minRooms and less than or equal to maxRooms.
*/
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

/*
FilterPrice returns a filter function that filters ImmoListings
based on their price. If minPrice is 0, it filters listings
with a price less than or equal to maxPrice. If maxPrice is 0,
it filters listings with a price greater than or equal to
minPrice. If both minPrice and maxPrice are 0, it returns a
filter function that always returns true. If minPrice and
maxPrice are both non-zero, it filters listings with a price
greater than or equal to minPrice and less than or equal to
maxPrice.
*/
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

/*
FilterArea returns a filter function that filters ImmoListings
based on their area. If minArea is 0, it filters listings
with an area less than or equal to maxArea. If maxArea is 0,
it filters listings with an area greater than or equal to
minArea. If both minArea and maxArea are 0, it returns a
filter function that always returns true. If minArea and
maxArea are both non-zero, it filters listings with an area
greater than or equal to minArea and less than or equal to
maxArea.
*/
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

/*
FilterKeyWords returns a filter function that filters ImmoListings
based on their description containing certain key words. The filter
function takes a slice of strings and returns true if any of the
strings are found in the description of the ImmoListing in a
case-insensitive manner. If the slice of strings is empty, it
returns a filter function that always returns true.
*/
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
