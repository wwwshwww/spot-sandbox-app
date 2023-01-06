package model

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"

type SpotsProfile struct {
	Key      int               `json:"key"`
	SpotKeys []spot.Identifier `json:"spotKeys"`
}
