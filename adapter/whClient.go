package adapter

import (
	"github.com/ehganzlieb/2025-10-26_willfahren/dto"
	whclient "github.com/ehganzlieb/2025-10-26_willfahren/whClient"
)

/*
 */
func WHClientDtoAdapter(wha *whclient.WHAdvert) *dto.Apartment {

	district, err := dto.DistrictFromPostCode(int(*wha.Postcode))
	if err != nil {
		//TODO: fault tolerance
		panic(err)
	}
	return &dto.Apartment{
		ID:          wha.ID,
		Title:       wha.Title,
		Description: wha.Description,
		Area:        float32(*wha.Area),
		Rooms:       0, //TODO: parse in whClient
		Price:       float32(*wha.Rent),
		District:    district,
		Location:    wha.Coordinates,
		URL:         *wha.URL,
	}
}
