package model

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
)

type ClusterElement struct {
	Key             int                       `json:"key"`
	DbscanProfileID dbscan_profile.Identifier `json:"dbscanProfileId"`
	SpotID          spot.Identifier           `json:"spotId"`
	AssignedNumber  int                       `json:"assignedNumber"`
	Paths           []*ClusterElement         `json:"paths"`
}
