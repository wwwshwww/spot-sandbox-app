package entity

import (
	"errors"
	"regexp"

	"github.com/wwwwshwww/spot-sandbox/external/yahoo"
)

type Address struct {
	PostalCode *string
	Prefecture *string
	City       *string
	Street1    *string
	Street2    *string
	Lat        *float64
	Lng        *float64
}

func New(
	postalCode, prefecture, city, street1, street2 *string,
) (Address, error) {
	var address = Address{
		PostalCode: postalCode,
		Prefecture: prefecture,
		City:       city,
		Street1:    street1,
		Street2:    street2,
	}
	lat, lng, err := address.CorrectLatLng()
	if err != nil {
		return Address{}, err
	}
	address.Lat = lat
	address.Lng = lng

	return address, nil
}

func NewAddressFromYahooReverseGeoCode(r *yahoo.ReverseGeoCodeResponse) (*Address, error) {
	if r.Feature[0].Property.Country.Code != "JP" {
		return nil, errors.New("日本以外は対応してないよ！")
	}
	target := r.Feature[0]

	prefecture := target.Property.AddressElement[0].Name
	city := target.Property.AddressElement[1].Name
	street1 := target.Property.AddressElement[2].Name

	pattern := regexp.MustCompile(street1)
	split := pattern.Split(target.Property.Address, 2)
	street2 := split[1]

	address, err := New(nil, &prefecture, &city, &street1, &street2)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (e *Address) CorrectLatLng() (*float64, *float64, error) {
	// TODO: implement correction latlng using geocoding
	return nil, nil, nil
}
