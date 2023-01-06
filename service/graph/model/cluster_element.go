package model

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
)

type ClusterElement struct {
	Key             int                       `json:"key"`
	DbscanProfileID dbscan_profile.Identifier `json:"dbscanProfileId"`
	SpotsProfileID  spots_profile.Identifier  `json:"spotsProfileId"`
	SpotID          spot.Identifier           `json:"spotId"`
	AssignedNumber  int                       `json:"assignedNumber"`
	Paths           []*ClusterElement         `json:"paths"`
}
