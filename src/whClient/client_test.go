package whclient

import (
	"testing"
)

func TestBullshit(t *testing.T) {
	/*
		d1, err := dto.DistrictByNumber(1)
		if err != nil {
			t.Fatal(err)
		}

		d2, err := dto.DistrictByNumber(2)
		if err != nil {
			t.Fatal(err)
		}
		d10, err := dto.DistrictByNumber(10)
		if err != nil {
			t.Fatal(err)
		}
	*/
	q := Query{
		MinPrice:     toPointerType(int64(500)),
		MaxPrice:     nil,
		MinArea:      nil,
		MaxArea:      toPointerType(int16(200)),
		Rooms1:       true,
		Rooms2:       true,
		Rooms3:       true,
		Rooms4:       false,
		Rooms5:       false,
		Rooms6to9:    false,
		Rooms10:      false,
		RoomsUnknown: true,
		/*Districts: []dto.District{
			*d1, *d2, *d10,
		},*/
	}

	wham, err := q.ProcessAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("len(wham):", len(*wham))
	t.Logf("%#v", wham)

}
