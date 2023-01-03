package model

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"

type SpotsProfile struct {
	ID      string            `json:"id"`
	SpotIDs []spot.Identifier `json:"spotIds"`
}
