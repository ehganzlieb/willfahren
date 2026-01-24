package cache

import (
	"github.com/ehganzlieb/2025-10-26_willfahren/adapter"
	"github.com/ehganzlieb/2025-10-26_willfahren/dto"
	whclient "github.com/ehganzlieb/2025-10-26_willfahren/whClient"
)

// WHCache internally uses whclient but only dto externally
type WHCache struct {
	Adverts map[uint64]whclient.WHAdvert
}

var whCache = WHCache{Adverts: make(map[uint64]whclient.WHAdvert)}

func Ingest(wha []whclient.WHAdvert) {
	for _, v := range wha {
		whCache.Adverts[v.ID] = v
	}
}
func All() []dto.Apartment {
	apts := make([]dto.Apartment, len(whCache.Adverts))
	for i, v := range whCache.Adverts {
		apts[i] = *adapter.WHClientDtoAdapter(&v)
	}
	return apts
}
