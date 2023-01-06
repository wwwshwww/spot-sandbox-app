package spots_profile_finder

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile/spots_profile"

type Finder interface {
	Find() ([]spots_profile.Identifier, error)
}
