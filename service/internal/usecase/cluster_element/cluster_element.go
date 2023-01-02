package cluster_element

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/cluster_element"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain_service"
)

type ClusterElementUsecase interface {
	Calc(dbscan_profile.Identifier, spots_profile.Identifier) ([]cluster_element.ClusterElement, error)
}

func New(
	sr spot.Repository,
	dpr dbscan_profile.Repository,
	spr spots_profile.Repository,
	cs domain_service.ClusteringService,
) ClusterElementUsecase {
	return clusterElementUsecase{
		sr:  sr,
		dpr: dpr,
		spr: spr,
		cs:  cs,
	}
}

type clusterElementUsecase struct {
	sr  spot.Repository
	dpr dbscan_profile.Repository
	spr spots_profile.Repository
	cs  domain_service.ClusteringService
}

func (u clusterElementUsecase) Calc(
	dp dbscan_profile.Identifier,
	spi spots_profile.Identifier,
) (
	[]cluster_element.ClusterElement,
	error,
) {
	return nil, nil
}
