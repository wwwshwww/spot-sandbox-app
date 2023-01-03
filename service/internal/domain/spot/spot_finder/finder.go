package spot_finder

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"

type Finder interface {
	Find(FilteringOptions, SortingOptions) ([]spot.Identifier, error)
}
