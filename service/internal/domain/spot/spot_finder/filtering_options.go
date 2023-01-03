package spot_finder

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"

type FilteringOptions struct {
	SpotIdentifiers []spot.Identifier
	PostalCode      string
}
