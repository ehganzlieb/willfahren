package dto

import "net/url"

type Apartment struct {
	Title       string
	Description string
	Area        float32
	Rooms       int
	Price       float32
	District    District
	Location    Coordinates
	StopsNearby []Stop
	URL         url.URL
}
