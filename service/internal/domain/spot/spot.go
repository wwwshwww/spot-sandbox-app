package spot

import (
	"errors"

	"github.com/wwwwshwww/spot-sandbox/internal/domain/address"
)

var (
	ErrNotFound       = errors.New("spot not found")
	defaultPreference = SpotPreferences{
		PostalCode:            "097-0000",
		AddressRepresentation: "稚内",
		Lat:                   45.4156,
		Lng:                   141.6734,
	}
)

type Spot interface {
	Identifier() Identifier
	Address() address.Address

	Overwrite(SpotPreferences) error
}

func New(i Identifier) Spot {
	address, err := address.New(
		defaultPreference.PostalCode,
		defaultPreference.AddressRepresentation,
		defaultPreference.Lat,
		defaultPreference.Lng,
	)
	if err != nil {
		panic("new spot error")
	}
	return &spot{
		identifier: i,
		address:    address,
	}
}

func Restore(
	i Identifier,
	sp SpotPreferences,
) Spot {
	s := New(i)
	if err := s.Overwrite(sp); err != nil {
		panic("restore spot error")
	}
	return s
}

type SpotPreferences struct {
	PostalCode            string
	AddressRepresentation string
	Lat                   float64
	Lng                   float64
}

type spot struct {
	identifier Identifier
	address    address.Address
}

func (e *spot) Identifier() Identifier   { return e.identifier }
func (e *spot) Address() address.Address { return e.address }

func (e *spot) Overwrite(p SpotPreferences) error {
	address, err := address.New(
		p.PostalCode,
		p.AddressRepresentation,
		p.Lat,
		p.Lng,
	)
	if err != nil {
		return err
	}
	e.address = address
	return nil
}
