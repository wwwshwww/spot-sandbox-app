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
	s, err := u.sr.Get(si)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u spotUsecase) Save(
	si spot.Identifier,
	sp spot.SpotPreferences,
) error {
	s, err := u.sr.Get(si)
	if err != nil {
		return err
	}
	if s == nil {
		s = spot.New(si)
	}
	if err := s.Overwrite(sp); err != nil {
		return err
	}
	if err := u.sr.Save(s); err != nil {
		return err
	}
	return nil
}
