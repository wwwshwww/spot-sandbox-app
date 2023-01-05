package spots_profile_graph

import (
	"strconv"

	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
)

func Marshal(sp spots_profile.SpotsProfile) *model.SpotsProfile {
	return &model.SpotsProfile{
		ID:      strconv.Itoa(int(sp.Identifier())),
		SpotIDs: sp.Spots(),
	}
}

func UnmarshalIdentifier(m string) spots_profile.Identifier {
	i, err := strconv.Atoi(m)
	if err != nil {
		panic("unmarshal spotsProfile identifier")
	}
	return spots_profile.Identifier(i)
}

func UnmarshalPreferences(m model.NewSpotsProfile) spots_profile.SpotsProfilePreferences {
	return spots_profile.SpotsProfilePreferences{
		Spots: common.Map(UnmarshalSpotIdentifier, m.SpotIds),
	}
}

func UnmarshalSpotIdentifier(m string) spot.Identifier {
	i, err := strconv.Atoi(m)
	if err != nil {
		panic("unmarshal spot identifier")
	}
	return spot.Identifier(i)
}
