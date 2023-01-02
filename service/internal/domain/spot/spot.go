package spot

import (
	"errors"

	"github.com/wwwwshwww/spot-sandbox/internal/domain/address"
)

var (
	ErrNotFound = errors.New("spot not found")
)

type Spot interface {
	Identifier() Identifier
	Address() address.Address
}

func New(i Identifier, a address.Address) Spot {
	return &spot{
		identifier: i,
		address:    a,
	}
}

func Restore(
	i Identifier,
	sp SpotPreferences,
) Spot {
	address, err := address.New(
		sp.PostalCode,
		sp.AddressRepresentation,
		sp.Lat,
		sp.Lng,
	)
	if err != nil {
		panic("spot restore error")
	}
	return New(i, address)
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
