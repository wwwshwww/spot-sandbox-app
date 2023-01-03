package spot_graph

import (
	"strconv"

	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
)

func Marshal(s spot.Spot) *model.Spot {
	return &model.Spot{
		ID:          strconv.Itoa(int(s.Identifier())),
		PostalCode:  s.Address().PostalCode(),
		AddressRepr: s.Address().AddressRepresentation(),
		Lat:         s.Address().Lat(),
		Lng:         s.Address().Lng(),
	}
}
