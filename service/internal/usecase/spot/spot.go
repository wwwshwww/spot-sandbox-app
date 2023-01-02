package spot

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot"

type SpotUsecase interface {
	Get(spot.Identifier) (spot.Spot, error)
	Save(spot.Identifier, spot.SpotPreferences) error
}

func New(sr spot.Repository) SpotUsecase {
	return spotUsecase{sr: sr}
}

type spotUsecase struct {
	sr spot.Repository
}

func (u spotUsecase) Get(si spot.Identifier) (
	spot.Spot,
	error,
) {
	return nil, nil
}

func (u spotUsecase) Save(
	si spot.Identifier,
	sp spot.SpotPreferences,
) error {
	return nil
}
