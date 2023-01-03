package spot

import (
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/google_maps"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot_finder"
)

type SpotUsecase interface {
	Get(spot.Identifier) (spot.Spot, error)
	BulkGet([]spot.Identifier) (map[spot.Identifier]spot.Spot, error)
	Save(spot.Identifier, common.LatLng) error

	ListAllSpots(spot_finder.FilteringOptions) ([]spot.Identifier, error)
}

func New(
	sr spot.Repository,
	sf spot_finder.Finder,
	gm *google_maps.GoogleMapsClient,
) SpotUsecase {
	return spotUsecase{
		sr: sr,
		sf: sf,
		gm: gm,
	}
}

type spotUsecase struct {
	sr spot.Repository
	sf spot_finder.Finder
	gm *google_maps.GoogleMapsClient
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

func (u spotUsecase) BulkGet(sis []spot.Identifier) (
	map[spot.Identifier]spot.Spot,
	error,
) {
	ss, err := u.sr.BulkGet(sis)
	if err != nil {
		return nil, err
	}
	return ss, nil
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

func (u spotUsecase) ListAllSpots(fo spot_finder.FilteringOptions) (
	[]spot.Identifier,
	error,
) {
	sis, err := u.sf.Find(fo, spot_finder.SortingOptions{
		Key:        spot_finder.SpotIdentifier,
		Descending: false,
	})
	if err != nil {
		return nil, err
	}
	return sis, nil
}
