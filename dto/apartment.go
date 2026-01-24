package dto

import "net/url"

type Apartment struct {
	ID          uint64
	Title       string
	Description string
	Area        float32
	Rooms       float32
	Price       float32
	District    *District
	Location    *Coordinates
	URL         url.URL
}
