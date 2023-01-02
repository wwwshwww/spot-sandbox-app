package cluster_element

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/cluster_element"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
)

type ClusterElementUsecase interface {
	Calc(dbscan_profile.Identifier, spots_profile.Identifier) ([]cluster_element.ClusterElement, error)
}
