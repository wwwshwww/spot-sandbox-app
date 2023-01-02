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

type SpotProfile interface {
	Identifier() Identifier
	Spots() []spot.Identifier

	UpdateSpots([]spot.Identifier) error
}

func New(i Identifier) SpotProfile {
	return &spotProfile{
		identifier: i,
		spots:      []spot.Identifier{},
	}
}

func Restore(i Identifier, spots []spot.Identifier) SpotProfile {
	s := New(i)
	if err := s.UpdateSpots(spots); err != nil {
		panic("resore panic")
	}
	return s
}

type spotProfile struct {
	identifier Identifier
	spots      []spot.Identifier
}

func (e *spotProfile) Identifier() Identifier   { return e.identifier }
func (e *spotProfile) Spots() []spot.Identifier { return e.spots }

func (e *spotProfile) UpdateSpots(s []spot.Identifier) error {
	if len(s) > spotMaxCount {
		return ErrSpotMaxCount
	}
	e.spots = s
	return nil
}

func (e *spotProfile) AppendSpot(s spot.Identifier) error {
	if len(e.spots)+1 > spotMaxCount {
		return ErrSpotMaxCount
	}
	if mapset.NewSet(e.spots...).Contains(s) {
		return ErrDuplicateSpot
	}

	e.spots = append(e.spots, s)
	return nil
}
