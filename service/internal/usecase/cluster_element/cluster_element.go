package cluster_element

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/cluster_element"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile/spots_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain_service"
)

// TODO: ClusterElementは値オブジェクトに近い存在なので集約ルートのClusterでusecaseを作成すべき

type ClusterElementUsecase interface {
	Calc(dbscan_profile.Identifier, []spot.Identifier) ([]cluster_element.ClusterElement, error)
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
	dpi dbscan_profile.Identifier,
	sis []spot.Identifier,
) (
	[]cluster_element.ClusterElement,
	error,
) {
	spots, err := u.sr.BulkGet(sis)
	if err != nil {
		return nil, err
	}
	dp, err := u.dpr.Get(dpi)
	if err != nil {
		return nil, err
	}
	ces, err := u.cs.DBScan(spots, dp)
	if err != nil {
		return nil, err
	}

	return ces, nil
}
