package spot_mysql

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot"

func unmarshal(row Spot) spot.Spot {
	return spot.Restore(
		spot.Identifier(row.ID),
		row.PostalCode,
		row.AddressRepresentation,
		row.Lat,
		row.Lng,
	)
}

func marshal(s spot.Spot) Spot {
	a := s.Address()
	return Spot{
		ID:                    s.Identifier().Value(),
		PostalCode:            a.PostalCode(),
		AddressRepresentation: a.AddressRepresentation(),
		Lat:                   a.Lat(),
		Lng:                   a.Lng(),
	}
}
