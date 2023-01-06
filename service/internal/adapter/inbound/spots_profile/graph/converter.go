package spots_profile_graph

import (
	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
)

func Marshal(sp spots_profile.SpotsProfile) *model.SpotsProfile {
	return &model.SpotsProfile{
		Key:      int(sp.Identifier()),
		SpotKeys: sp.Spots(),
	}
}

func UnmarshalIdentifier(m int) spots_profile.Identifier {
	return spots_profile.Identifier(m)
}

func UnmarshalPreferences(m model.NewSpotsProfile) spots_profile.SpotsProfilePreferences {
	return spots_profile.SpotsProfilePreferences{
		Spots: common.Map(UnmarshalSpotIdentifier, m.SpotKeys),
	}
}

func UnmarshalSpotIdentifier(m int) spot.Identifier {
	return spot.Identifier(m)
}
