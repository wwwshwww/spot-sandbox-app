package dbscan_profile_finder

import "github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"

type Finder interface {
	Find() ([]dbscan_profile.Identifier, error)
}
