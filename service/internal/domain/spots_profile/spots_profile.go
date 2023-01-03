package spots_profile

import (
	"errors"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot"
)

var (
	spotMaxCount     = 2000
	ErrSpotMaxCount  = errors.New("exceed max count")
	ErrDuplicateSpot = errors.New("duplication spot")
)

type SpotsProfile interface {
	Identifier() Identifier
	Spots() []spot.Identifier

	UpdateSpots([]spot.Identifier) error
	AppendSpot(spot.Identifier) error
	RemoveSpot(spot.Identifier) error

	Overwrite(SpotsProfilePreferences) error
}

func New(i Identifier) SpotsProfile {
	return &spotsProfile{
		identifier: i,
		spots:      []spot.Identifier{},
	}
}

func Restore(i Identifier, spp SpotsProfilePreferences) SpotsProfile {
	s := New(i)
	if err := s.UpdateSpots(spp.Spots); err != nil {
		panic("resore panic")
	}
	return s
}

type SpotsProfilePreferences struct {
	Spots []spot.Identifier
}

type spotsProfile struct {
	identifier Identifier
	spots      []spot.Identifier
}

func (e *spotsProfile) Identifier() Identifier   { return e.identifier }
func (e *spotsProfile) Spots() []spot.Identifier { return e.spots }

func (e *spotsProfile) UpdateSpots(s []spot.Identifier) error {
	if len(s) > spotMaxCount {
		return ErrSpotMaxCount
	}
	e.spots = mapset.NewSet(s...).ToSlice()
	return nil
}

func (e *spotsProfile) AppendSpot(s spot.Identifier) error {
	if err := e.UpdateSpots(append(e.spots, s)); err != nil {
		return err
	}
	return nil
}

func (e *spotsProfile) RemoveSpot(s spot.Identifier) error {
	newSpots := make([]spot.Identifier, 0, len(e.spots))
	for _, si := range e.spots {
		if s != si {
			newSpots = append(newSpots, si)
		}
	}
	if err := e.UpdateSpots(newSpots); err != nil {
		return err
	}
	return nil
}

func (e *spotsProfile) Overwrite(spp SpotsProfilePreferences) error {
	if err := e.UpdateSpots(spp.Spots); err != nil {
		return err
	}
	return nil
}
