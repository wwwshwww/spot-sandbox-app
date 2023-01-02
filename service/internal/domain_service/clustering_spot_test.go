package domain_service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain_service"
)

func TestDBScan(t *testing.T) {
	ctx := context.Background()
	cs := domain_service.NewClusteringService(ctx, nil, nil, nil)

	dbscanProfiles := []dbscan_profile.DbscanProfile{
		dbscan_profile.Restore(
			dbscan_profile.Identifier(1),
			dbscan_profile.DbscanProfilePreferences{
				DistanceType:   dbscan_profile.Hubeny,
				MinCount:       0,
				MeterThreshold: func(n int) *int { return &n }(10000),
			},
		),
	}

	spots := []spot.Spot{
		spot.Restore(
			spot.Identifier(1),
			"000-0001",
			"わりと場所1-1",
			0.0,
			0.0,
		),
		spot.Restore(
			spot.Identifier(2),
			"000-0002",
			"そこそこ近い場所2-2",
			0.01,
			-0.01,
		),
		spot.Restore(
			spot.Identifier(3),
			"000-0003",
			"まあまあ近い場所3-3",
			0.02,
			-0.02,
		),
		spot.Restore(
			spot.Identifier(4),
			"000-0003",
			"たぶん近い場所4-4",
			0.018,
			-0.018,
		),
		spot.Restore(
			spot.Identifier(5),
			"000-0003",
			"ちょっと近い場所5-5",
			0.019,
			-0.019,
		),
		spot.Restore(
			spot.Identifier(6),
			"000-0006",
			"ふつうに遠い場所6-6",
			0.16,
			0.16,
		),
	}

	spotMap := make(map[spot.Identifier]spot.Spot, len(spots))
	for _, s := range spots {
		spotMap[s.Identifier()] = s
	}

	spotProfiles := []spots_profile.SpotProfile{
		spots_profile.Restore(
			spots_profile.Identifier(1),
			common.Map(func(s spot.Spot) spot.Identifier { return s.Identifier() }, spots),
		),
	}

	tests := []struct {
		dp             dbscan_profile.DbscanProfile
		sp             spots_profile.SpotProfile
		expectCountMap map[int]int
	}{
		{
			dbscanProfiles[0],
			spotProfiles[0],
			map[int]int{
				1: 1,
				2: 5,
			},
		},
	}

	for _, tt := range tests {
		ces, err := cs.DBScan(spotMap, tt.dp, tt.sp)
		assert.NoError(t, err)

		countMap := make(map[int]int)
		for _, ce := range ces {
			if _, ok := countMap[ce.AssignedNumber()]; !ok {
				countMap[ce.AssignedNumber()] = 1
			} else {
				countMap[ce.AssignedNumber()] += 1
			}
		}

		assert.Equal(t, len(tt.expectCountMap), len(countMap))
		for k, expected := range tt.expectCountMap {
			actual, ok := countMap[k]
			assert.True(t, ok)
			assert.Equal(t, expected, actual)
		}
	}
}
