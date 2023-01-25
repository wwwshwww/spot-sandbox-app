package domain_service

import (
	"context"
	"errors"
	"math"
	"sort"
	"time"

	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/cache"
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/google_maps"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/cluster_element"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
)

var (
	majorRadius  = 6378137.0
	minorRadius  = 6356752.314245
	eccentricity = math.Sqrt((math.Pow(majorRadius, 2) - math.Pow(minorRadius, 2)) / math.Pow(majorRadius, 2))
)

type ClusteringService interface {
	DBScan(
		map[spot.Identifier]spot.Spot,
		dbscan_profile.DbscanProfile,
	) (
		int,
		[]cluster_element.ClusterElement,
		error,
	)
}

type clusteringService struct {
	ctx                 context.Context
	googleMapsClient    *google_maps.GoogleMapsClient
	distanceCacheClient *cache.DistanceCacheClient
	durationCacheClient *cache.DurationCacheClient
}

func NewClusteringService(
	ctz context.Context,
	gmc *google_maps.GoogleMapsClient,
	dicc *cache.DistanceCacheClient,
	ducc *cache.DurationCacheClient,
) ClusteringService {
	return &clusteringService{
		ctx:                 ctz,
		googleMapsClient:    gmc,
		distanceCacheClient: dicc,
		durationCacheClient: ducc,
	}
}

// TODO: cluster集約のメソッドとして書けそう
func (c *clusteringService) DBScan(
	spots map[spot.Identifier]spot.Spot,
	dbscanProfile dbscan_profile.DbscanProfile,
) (
	int,
	[]cluster_element.ClusterElement,
	error,
) {
	if dbscanProfile == nil {
		return 0, nil, errors.New("dbscan profile not found")
	}
	if len(spots) == 0 {
		return 0, nil, errors.New("zero spot")
	}

	// 対象地点を緯度で大きい順にソートする
	targetSpotIDs := make([]spot.Identifier, 0, len(spots))
	for si := range spots {
		targetSpotIDs = append(targetSpotIDs, si)
	}
	sort.SliceStable(targetSpotIDs, func(i, j int) bool {
		return spots[targetSpotIDs[i]].Address().Lat() > spots[targetSpotIDs[j]].Address().Lat()
	})

	// 結果として返すclusterElementの原型を作成
	clusterElements := make([]cluster_element.ClusterElement, len(spots))
	for i, si := range targetSpotIDs {
		clusterElements[i] = cluster_element.New(
			cluster_element.Identifier(i),
			dbscanProfile.Identifier(),
			si,
		)
	}

	// 各種Mapを作成しておく
	cei2ceMap := make(map[cluster_element.Identifier]cluster_element.ClusterElement, len(clusterElements))
	cei2siMap := make(map[cluster_element.Identifier]spot.Identifier, len(clusterElements))
	si2ceiMap := make(map[spot.Identifier]cluster_element.Identifier, len(clusterElements))
	for _, ce := range clusterElements {
		cei2ceMap[ce.Identifier()] = ce
		cei2siMap[ce.Identifier()] = ce.SpotIdentifier()
		si2ceiMap[ce.SpotIdentifier()] = ce.Identifier()
	}

	// 各地点ごとの近傍ノードを算出
	var pathMap map[spot.Identifier][]spot.Identifier
	var err error
	switch dbscanProfile.DistanceType() {
	case dbscan_profile.Hubeny:
		pathMap, err = c.GetPathMapWithInt(
			spots,
			targetSpotIDs,
			*dbscanProfile.MeterThreshold(),
			c.GetHubenys,
		)
	case dbscan_profile.RouteLength:
		pathMap, err = c.GetPathMapWithInt(
			spots,
			targetSpotIDs,
			*dbscanProfile.MeterThreshold(),
			c.GetRouteLengths,
		)
	case dbscan_profile.TravelTime:
		pathMap, err = c.GetPathMapWithDuration(
			spots,
			targetSpotIDs,
			*dbscanProfile.DurationThreshold(),
			c.GetTravelTimes,
		)
	}
	if err != nil {
		return 0, nil, err
	}

	/*
		DBScanによるクラスタリングを行い、クラスタを記録していく。番号の意味は以下の通り。
		 - 0:		未処理
		 - -1:		最低個数に満たないクラスタ
		 - other:	クラスタ番号
	*/
	assigns := make(map[int][]cluster_element.Identifier)
	limit := dbscanProfile.MaxCount()
	isFull := func(l []cluster_element.Identifier) bool {
		return limit != nil && len(l) == int(*limit)
	}

	clsNum := 1
	for _, ce := range clusterElements {
		if ce.IsAssigned() {
			continue
		}
		assigns[clsNum] = make([]cluster_element.Identifier, 0, len(clusterElements))

		q := common.NewDeque[cluster_element.Identifier]()
		q.AppendLeft(ce.Identifier())
		for q.Len() > 0 {
			cei := q.PopLeft()
			if isFull(assigns[clsNum]) {
				continue
			}
			si := cei2siMap[cei]

			// clsNumのクラスタに所属させる
			assigns[clsNum] = append(assigns[clsNum], cei)
			if err := cei2ceMap[cei].UpdateAssignedNumber(clsNum); err != nil {
				return 0, nil, err
			}
			if err := cei2ceMap[cei].UpdatePaths(common.Map(
				func(i spot.Identifier) cluster_element.Identifier {
					return si2ceiMap[i]
				},
				pathMap[si],
			)); err != nil {
				return 0, nil, err
			}

			// 近傍ノードの探索
			for _, nsi := range pathMap[si] {
				ncei := si2ceiMap[nsi]
				if cei2ceMap[ncei].IsAssigned() {
					continue
				}
				// 双方向で近傍とみなせる場合のみ所属候補とする
				if common.Contain(pathMap[nsi], si) {
					q.Append(ncei)
				}
			}
		}

		if len(assigns[clsNum]) < int(dbscanProfile.MinCount()) {
			// クラスタの要素数が最小個数に満たない場合はクラスタとみなさない
			for _, cei := range assigns[clsNum] {
				if err := cei2ceMap[cei].LackAssign(); err != nil {
					return 0, nil, err
				}
			}
		} else {
			// クラスタの要素数が最小個数を満たしている場合はおk
			clsNum++
		}
	}

	clusterCount := clsNum
	if clusterElements[len(clusterElements)-1].AssignedNumber() != cluster_element.CLUSTER_LACK {
		clusterCount -= 1
	}
	return clusterCount, clusterElements, nil
}

func (c *clusteringService) GetHubenys(
	from common.LatLng,
	tos []common.LatLng,
) (
	[]int,
	error,
) {
	fLat := from.Lat * (math.Pi / 180.0)
	fLng := from.Lng * (math.Pi / 180.0)
	ds := make([]int, len(tos))
	for i, to := range tos {
		tLat := to.Lat * (math.Pi / 180.0)
		tLng := to.Lng * (math.Pi / 180.0)

		dy := fLat - tLat
		dx := fLng - tLng
		p := (fLat + tLat) / 2
		w := math.Sqrt(1 - math.Pow(eccentricity, 2)*math.Pow(math.Sin(p), 2))
		m := (majorRadius * (1 - math.Pow(eccentricity, 2))) / math.Pow(w, 3)
		n := majorRadius / w
		ds[i] = int(math.Abs(math.Sqrt(math.Pow(dy*m, 2) + math.Pow(dx*n*math.Cos(p), 2))))
	}
	return ds, nil
}

func (c *clusteringService) GetRouteLengths(
	from common.LatLng,
	tos []common.LatLng,
) (
	[]int,
	error,
) {
	// キャッシュから距離値の取得を試みる
	fromTos := make([]cache.FromToLatLng, len(tos))
	for i, to := range tos {
		fromTos[i] = cache.FromToLatLng{
			From: from,
			To:   to,
		}
	}
	dist, nfIndex, err := c.distanceCacheClient.BulkGet(c.ctx, fromTos)
	if err != nil {
		return nil, err
	}

	// キャッシュヒットしなかった箇所がある場合は補完する
	if len(nfIndex) > 0 {
		// 補完する分をGoogleMapsから取得する
		tosN := make([]common.LatLng, len(nfIndex))
		for i, nf := range nfIndex {
			tosN[i] = tos[nf]
		}
		resDura, resDist, err := c.googleMapsClient.GetDurationAndDistanceOneToMany(from, tosN)
		if err != nil {
			return nil, err
		}

		// さきほどキャッシュヒットしなかった箇所を穴埋めする
		for j, nf := range nfIndex {
			dist[nf] = &resDist[j]
		}

		// ついでに新しくキャッシュしておく
		fromTos := make([]cache.FromToLatLng, len(nfIndex))
		for j, toN := range tosN {
			fromTos[j] = cache.FromToLatLng{
				From: from,
				To:   toN,
			}
		}
		err = c.durationCacheClient.BulkSet(c.ctx, fromTos, resDura)
		if err != nil {
			return nil, err
		}
		err = c.distanceCacheClient.BulkSet(c.ctx, fromTos, resDist)
		if err != nil {
			return nil, err
		}
	}

	return common.Map(func(n *int) int { return *n }, dist), nil
}

func (c *clusteringService) GetTravelTimes(
	from common.LatLng,
	tos []common.LatLng,
) (
	[]time.Duration,
	error,
) {
	// キャッシュから距離値の取得を試みる
	fromTos := make([]cache.FromToLatLng, len(tos))
	for i, to := range tos {
		fromTos[i] = cache.FromToLatLng{
			From: from,
			To:   to,
		}
	}
	dura, nfIndex, err := c.durationCacheClient.BulkGet(c.ctx, fromTos)
	if err != nil {
		return nil, err
	}

	// キャッシュヒットしなかった箇所がある場合は補完する
	if len(nfIndex) > 0 {
		// 補完する分をGoogleMapsから取得する
		tosN := make([]common.LatLng, len(nfIndex))
		for i, nf := range nfIndex {
			tosN[i] = tos[nf]
		}
		resDura, resDist, err := c.googleMapsClient.GetDurationAndDistanceOneToMany(from, tosN)
		if err != nil {
			return nil, err
		}

		// さきほどキャッシュヒットしなかった箇所を穴埋めする
		for j, nf := range nfIndex {
			dura[nf] = &resDura[j]
		}

		// ついでに新しくキャッシュしておく
		fromTos := make([]cache.FromToLatLng, len(nfIndex))
		for j, toN := range tosN {
			fromTos[j] = cache.FromToLatLng{
				From: from,
				To:   toN,
			}
		}
		err = c.durationCacheClient.BulkSet(c.ctx, fromTos, resDura)
		if err != nil {
			return nil, err
		}
		err = c.distanceCacheClient.BulkSet(c.ctx, fromTos, resDist)
		if err != nil {
			return nil, err
		}
	}

	return common.Map(func(n *time.Duration) time.Duration { return *n }, dura), nil
}

// 近傍ノードを近い順にソート済みの状態で取得する。閾値としてIntを扱うバージョン
func (c *clusteringService) GetPathMapWithInt(
	spots map[spot.Identifier]spot.Spot,
	spotIDs []spot.Identifier,
	threshold int,
	distFn func(origin common.LatLng, dests []common.LatLng) ([]int, error),
) (map[spot.Identifier][]spot.Identifier, error) {
	latLngMap := c.GetLatLngMap(spots, spotIDs)

	toLatLngs := make([]common.LatLng, len(spotIDs))
	for i, dstSpot := range spotIDs {
		toLatLngs[i] = latLngMap[dstSpot]
	}

	pathMap := make(map[spot.Identifier][]spot.Identifier)
	for _, oriSpot := range spotIDs {
		distances, err := distFn(latLngMap[oriSpot], toLatLngs)
		if err != nil {
			return nil, err
		}

		paths := make([]spot.Identifier, 0, len(spotIDs))
		spotToDist := make(map[spot.Identifier]int)
		for j, si := range spotIDs {
			if si != oriSpot && distances[j] < threshold {
				spotToDist[si] = distances[j]
				paths = append(paths, si)
			}
		}
		sort.Slice(paths, func(i, j int) bool {
			return spotToDist[paths[i]] < spotToDist[paths[j]]
		})
		pathMap[oriSpot] = paths
	}

	return pathMap, nil
}

// 近傍ノードを近い順にソート済みの状態で取得する。閾値としてtime.Durationを扱うバージョン
func (c *clusteringService) GetPathMapWithDuration(
	spots map[spot.Identifier]spot.Spot,
	spotIDs []spot.Identifier,
	threshold time.Duration,
	distFn func(origin common.LatLng, dests []common.LatLng) ([]time.Duration, error),
) (map[spot.Identifier][]spot.Identifier, error) {
	latLngMap := c.GetLatLngMap(spots, spotIDs)

	toLatLngs := make([]common.LatLng, len(spotIDs))
	for i, dstSpot := range spotIDs {
		toLatLngs[i] = latLngMap[dstSpot]
	}

	pathMap := make(map[spot.Identifier][]spot.Identifier)
	for _, oriSpot := range spotIDs {
		distances, err := distFn(latLngMap[oriSpot], toLatLngs)
		if err != nil {
			return nil, err
		}

		paths := make([]spot.Identifier, 0, len(spotIDs))
		spotToDist := make(map[spot.Identifier]time.Duration)
		for j, si := range spotIDs {
			if si != oriSpot && distances[j] < threshold {
				spotToDist[si] = distances[j]
				paths = append(paths, si)
			}
		}
		sort.Slice(paths, func(i, j int) bool {
			return spotToDist[paths[i]] < spotToDist[paths[j]]
		})
		pathMap[oriSpot] = paths
	}

	return pathMap, nil
}

func (c *clusteringService) GetLatLngMap(
	spots map[spot.Identifier]spot.Spot,
	spotIDs []spot.Identifier,
) map[spot.Identifier]common.LatLng {
	latLngMap := make(map[spot.Identifier]common.LatLng, len(spotIDs))
	for _, si := range spotIDs {
		latLngMap[si] = common.LatLng{
			Lat: spots[si].Address().Lat(),
			Lng: spots[si].Address().Lng(),
		}
	}
	return latLngMap
}
