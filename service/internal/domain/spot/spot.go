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
	postalCode, addressRepresentation string,
	lat, lng float64,
) Spot {
	address, err := address.New(
		postalCode,
		addressRepresentation,
		lat,
		lng,
	)
	if err != nil {
		panic("spot restore error")
	}
	return New(i, address)
}

type spot struct {
	identifier Identifier
	address    address.Address
}

func (e *spot) Identifier() Identifier   { return e.identifier }
func (e *spot) Address() address.Address { return e.address }
