package whclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"maps"
	"net/url"
	"strconv"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/ehganzlieb/willwohnen/src/dto"
)

type Query struct {
	Districts    []dto.District
	MinPrice     *int64
	MaxPrice     *int64
	MinArea      *int16
	MaxArea      *int16
	Rooms1       bool
	Rooms2       bool
	Rooms3       bool
	Rooms4       bool
	Rooms5       bool
	Rooms6to9    bool
	Rooms10      bool
	RoomsUnknown bool
	page         *int64 //pagination
}

const WHImmoBaseURL = "https://www.willhaben.at/iad/immobilien/mietwohnungen/mietwohnung-angebote"
const MinPriceField = "PRICE_FROM"
const MaxPriceField = "PRICE_TO"
const MinAreaField = "ESTATE_SIZE/LIVING_AREA_FROM"
const MaxAreaField = "ESTATE_SIZE/LIVING_AREA_TO"
const RoomsField = "NO_OF_ROOMS_BUCKET"
const Rooms1 = "1X1"
const Rooms2 = "2X2"
const Rooms3 = "3X3"
const Rooms4 = "4X4"
const Rooms5 = "5X5"
const Rooms6to9 = "6X9"
const Rooms10plus = "10X"
const RoomsUnknown = "0X0"

type WHAdvert struct {
	ID           uint64
	Title        string
	Heading      string
	Postcode     *uint64
	LocationID   *uint64
	URL          *url.URL
	Description  string
	SellerName   string
	Area         *uint64
	Coordinates  *dto.Coordinates
	Rent         *float64
	PrivateOffer bool
	PublishTime  *time.Time
	Images       []url.URL
}

type WHAdvertMap map[uint64]WHAdvert

func (wham *WHAdvertMap) Merge(other WHAdvertMap) WHAdvertMap {
	maps.Copy((*wham), other)
	return *wham
}

type WHQueryResult struct {
	RowsInSet     int
	RowsTotal     int
	RowsRequested int
	Adverts       WHAdvertMap
}

func (q Query) URL() (*url.URL, error) {
	u, err := url.Parse(WHImmoBaseURL)
	if err != nil {
		return nil, err
	}
	uq := u.Query()
	if q.MinPrice != nil {
		uq.Set(MinPriceField, strconv.FormatInt(*q.MinPrice, 10))
	}
	if q.MaxPrice != nil {
		uq.Set(MaxPriceField, strconv.FormatInt(*q.MaxPrice, 10))
	}
	if q.MinArea != nil {
		uq.Set(MinAreaField, strconv.FormatInt(int64(*q.MinArea), 10))
	}
	if q.MaxArea != nil {
		uq.Set(MaxAreaField, strconv.FormatInt(int64(*q.MaxArea), 10))
	}
	if q.Rooms1 {
		uq.Set(RoomsField, Rooms1)
	}
	if q.Rooms2 {
		uq.Set(RoomsField, Rooms2)
	}
	if q.Rooms3 {
		uq.Set(RoomsField, Rooms3)
	}
	if q.Rooms4 {
		uq.Set(RoomsField, Rooms4)
	}
	if q.Rooms5 {
		uq.Set(RoomsField, Rooms5)
	}
	if q.Rooms6to9 {
		uq.Set(RoomsField, Rooms6to9)
	}
	if q.Rooms10 {
		uq.Set(RoomsField, Rooms10plus)
	}
	if q.RoomsUnknown {
		uq.Set(RoomsField, RoomsUnknown)
	}

	for _, d := range q.Districts {
		id, err := AreaID(&d)
		if err != nil {
			return nil, err
		}
		uq.Add("areaId", strconv.FormatUint(id, 10))
	}

	if q.page != nil {
		uq.Add("page", strconv.FormatInt(*q.page, 10))
	}

	u.RawQuery = uq.Encode()
	return u, nil

}

func (q Query) ProcessAll() (*WHAdvertMap, error) {
	//first process, then followup until all results are fetched
	whq, err := q.Process()
	if err != nil {
		return nil, err
	}

	wham := whq.Adverts
	rowsInSet := whq.RowsInSet
	rowsRequested := whq.RowsRequested
	for rowsInSet == rowsRequested {
		whq, err = whq.FollowUp(&q)
		if err != nil {
			return nil, err
		}
		wham = wham.Merge(whq.Adverts)
		rowsInSet = whq.RowsInSet
		rowsRequested = whq.RowsRequested
	}
	return &wham, nil
}

func (q Query) Process() (*WHQueryResult, error) {
	var whd WHQueryResult
	qu, err := q.URL()
	if err != nil {
		return nil, err
	}
	html, err := soup.Get(qu.String())
	if err != nil {
		return nil, err
	}

	root := soup.HTMLParse(html)
	whd, err = interpretWHData(root)
	if err != nil {
		return nil, err
	}

	return &whd, nil

}

// Request with incremented by one or set to two (if nil) page
func (whd *WHQueryResult) FollowUp(q *Query) (*WHQueryResult, error) {
	if q.page == nil {
		q.page = toPointerType(int64(2))
	} else {
		*q.page++
	}
	// add map elements to whd
	return q.Process()
}

func interpretWHData(r soup.Root) (WHQueryResult, error) {
	var whd WHQueryResult
	whJson := r.Find("script", "type", "application/json")
	var data interface{}
	json.Unmarshal([]byte(whJson.FullText()), &data)
	//t.Logf("%v", data)
	props := data.(map[string]interface{})["props"]
	pageProps := props.(map[string]interface{})["pageProps"]
	searchResult := pageProps.(map[string]interface{})["searchResult"]
	whd.RowsTotal = int(searchResult.(map[string]interface{})["rowsFound"].(float64))
	whd.RowsInSet = int(searchResult.(map[string]interface{})["rowsReturned"].(float64))
	whd.RowsRequested = int(searchResult.(map[string]interface{})["rowsRequested"].(float64))
	adverts := searchResult.(map[string]interface{})["advertSummaryList"].(map[string]interface{})["advertSummary"].([]interface{})
	advertMap := make(WHAdvertMap)
	for _, advert := range adverts {
		var err error
		if advertMap, err = advertMap.parseAdvert(advert.(map[string]interface{})); err != nil {
			log.Default().Println(err)
		}
	}
	whd.Adverts = advertMap
	return whd, nil
}

func (wam WHAdvertMap) parseAdvert(rawAd map[string]interface{}) (WHAdvertMap, error) {
	var adv WHAdvert
	idString := rawAd["id"].(string)
	log.Println("parsing advert ", idString)
	var err error
	if adv.ID, err = strconv.ParseUint(idString, 10, 64); err != nil {
		return wam, err
	}
	if rawAd["description"] == nil {
		return wam, errors.New("no title in listing")
	}
	attrs := rawAd["attributes"].(map[string]interface{})
	attrArr := attrs["attribute"].([]interface{})

	//[]map[string]interface{})["attribute"]

	adv.Title = rawAd["description"].(string)

	for _, _a := range attrArr {
		a := _a.(map[string]interface{})
		switch a["name"] {
		case "BODY_DYN":
			adv.Description = firstStringVal(a)
		case "ORGNAME":
			adv.SellerName = firstStringVal(a)
		case "HEADING":
			adv.Heading = firstStringVal(a)
		case "SEO_URL":
			u, err := url.Parse("https://willhaben.at/iad")
			if err != nil {
				log.Default().Println(err)
			} else {
				adv.URL = u.JoinPath(firstStringVal(a))
			}
		case "RENT/PER_MONTH_LETTINGS":
			f, err := strconv.ParseFloat(firstStringVal(a), 64)
			if err != nil {
				log.Println(err)
			} else {
				adv.Rent = &f
			}
		case "ISPRIVATE":
			switch firstStringVal(a) {
			case "1":
				adv.PrivateOffer = true
			default:
				adv.PrivateOffer = false
			}
		case "LOCATION_ID":
			u, err := strconv.ParseUint(firstStringVal(a), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				adv.LocationID = &u
			}
		case "POSTCODE":
			u, err := strconv.ParseUint(firstStringVal(a), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				adv.Postcode = &u
			}
		case "ESTATE_SIZE":
			u, err := strconv.ParseUint(firstStringVal(a), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				adv.Area = &u
			}
		case "PUBLISHED":
			i, err := strconv.ParseInt(firstStringVal(a), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				t := time.UnixMilli(i)
				adv.PublishTime = &t
			}
		case "COORDINATES":
			var f1, f2 float64
			_, err := fmt.Sscanf(firstStringVal(a), "%f,%f", &f1, &f2)
			if err != nil {
				log.Println(err)
			} else {
				adv.Coordinates = &dto.Coordinates{
					X: f1,
					Y: f2,
				}
			}
		}
	}
	wam[adv.ID] = adv
	return wam, nil
}

func firstStringVal(a map[string]interface{}) string {
	return a["values"].([]interface{})[0].(string)
}

func toPointerType[T any](t T) *T {
	return &t
}
