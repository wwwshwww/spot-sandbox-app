package spot

import (
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/google_maps"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot"
)

type SpotUsecase interface {
	Get(spot.Identifier) (spot.Spot, error)
	Save(spot.Identifier, common.LatLng) error
}

func New(sr spot.Repository) SpotUsecase {
	return spotUsecase{sr: sr}
}

type spotUsecase struct {
	sr spot.Repository
	gm google_maps.GoogleMapsClient
}

func (u spotUsecase) Get(si spot.Identifier) (
	spot.Spot,
	error,
) {
	s, err := u.sr.Get(si)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u spotUsecase) Save(
	si spot.Identifier,
	latlng common.LatLng,
) error {
	s, err := u.sr.Get(si)
	if err != nil {
		return err
	}
	if s == nil {
		s = spot.New(si)
	}
	postalCode, addrRepr, err := u.gm.ReverseGeocode(latlng.Lat, latlng.Lng)
	if err != nil {
		return err
	}
	if err := s.Overwrite(
		spot.SpotPreferences{
			PostalCode:            postalCode,
			AddressRepresentation: addrRepr,
			Lat:                   latlng.Lat,
			Lng:                   latlng.Lng,
		},
	); err != nil {
		return err
	}
	if err := u.sr.Save(s); err != nil {
		return err
	}
	return nil
}
