package cache

// WHCache internally uses whclient but only dto externally
type WHCache struct {
	Adverts map[uint64]whclient.WHAdvert
}

var whCache = WHCache{Adverts: make(map[uint64]whclient.WHAdvert)}

func Ingest(wha []WHAdvert) {
	for _, v := range wha {
		whCache.Adverts[v.ID] = v
	}
}
func All() []dto.Apartment {
	apts := make([]dto.Apartment, len(whCache.Adverts))
	i := 0
	for _, v := range whCache.Adverts {
		apts[i] = dto.Apartment{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Area:        v.Area,
			Rooms:       v.Rooms,
			Price:       v.Price,
			District:    v.District,
			Location:    v.Location,
			URL:         v.URL,
		}
		i++
	}
	return apts
}
