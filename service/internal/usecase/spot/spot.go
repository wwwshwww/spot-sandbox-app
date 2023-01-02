package spot

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot"

type SpotUsecase interface {
	Get(spot.Identifier) (spot.Spot, error)
	Save(spot.Identifier, spot.SpotPreferences) error
}
