package domain_service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/cluster_element"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile/spots_profile"
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
		dbscan_profile.Restore(
			dbscan_profile.Identifier(2),
			dbscan_profile.DbscanProfilePreferences{
				DistanceType:   dbscan_profile.Hubeny,
				MinCount:       0,
				MaxCount:       func(n int) *int { return &n }(3),
				MeterThreshold: func(n int) *int { return &n }(10000),
			},
		),
		dbscan_profile.Restore(
			dbscan_profile.Identifier(3),
			dbscan_profile.DbscanProfilePreferences{
				DistanceType:   dbscan_profile.Hubeny,
				MinCount:       2,
				MaxCount:       func(n int) *int { return &n }(2),
				MeterThreshold: func(n int) *int { return &n }(10000),
			},
		),
	}

	spots := []spot.Spot{
		spot.Restore(
			spot.Identifier(1),
			spot.SpotPreferences{
				PostalCode:            "000-0001",
				AddressRepresentation: "わりと場所1-1",
				Lat:                   0.0,
				Lng:                   0.0,
			},
		),
		spot.Restore(
			spot.Identifier(2),
			spot.SpotPreferences{
				PostalCode:            "000-0002",
				AddressRepresentation: "そこそこ近い場所2-2",
				Lat:                   0.01,
				Lng:                   -0.01,
			},
		),
		spot.Restore(
			spot.Identifier(3),
			spot.SpotPreferences{
				PostalCode:            "000-0003",
				AddressRepresentation: "まあまあ近い場所3-3",
				Lat:                   0.02,
				Lng:                   -0.02,
			},
		),
		spot.Restore(
			spot.Identifier(4),
			spot.SpotPreferences{
				PostalCode:            "000-0003",
				AddressRepresentation: "たぶん近い場所4-4",
				Lat:                   0.018,
				Lng:                   -0.018,
			},
		),
		spot.Restore(
			spot.Identifier(5),
			spot.SpotPreferences{
				PostalCode:            "000-0003",
				AddressRepresentation: "ちょっと近い場所5-5",
				Lat:                   0.019,
				Lng:                   -0.019,
			},
		),
		spot.Restore(
			spot.Identifier(6),
			spot.SpotPreferences{
				PostalCode:            "000-0006",
				AddressRepresentation: "1から25000mくらい離れてる場所6",
				Lat:                   0.16,
				Lng:                   0.16,
			},
		),
	}

	spotMap := make(map[spot.Identifier]spot.Spot, len(spots))
	for _, s := range spots {
		spotMap[s.Identifier()] = s
	}

	spotProfiles := []spots_profile.SpotsProfile{
		spots_profile.Restore(
			spots_profile.Identifier(1),
			spots_profile.SpotsProfilePreferences{
				Spots: common.Map(func(s spot.Spot) spot.Identifier { return s.Identifier() }, spots),
			},
		),
	}

	tests := []struct {
		desc           string
		dp             dbscan_profile.DbscanProfile
		sp             spots_profile.SpotsProfile
		expectCountMap map[int]int
	}{
		{
			"Hubeny: 1個,5個と分類されるケース",
			dbscanProfiles[0],
			spotProfiles[0],
			map[int]int{
				1: 1,
				2: 5,
			},
		},
		{
			"Hubeny: 1個,3個,2個と分類されるケース",
			dbscanProfiles[1],
			spotProfiles[0],
			map[int]int{
				1: 1,
				2: 3,
				3: 2,
			},
		},
		{
			"Hubeny: 1個,2個,2個,1個と分類されMinCount2に満たないやつが2つ出るケース",
			dbscanProfiles[2],
			spotProfiles[0],
			map[int]int{
				cluster_element.CLUSTER_LACK: 2,
				1:                            2,
				2:                            2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			spots := make(map[spot.Identifier]spot.Spot)
			for _, si := range tt.sp.Spots() {
				spots[si] = spotMap[si]
			}
			cnum, ces, err := cs.DBScan(spots, tt.dp)
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
			assert.Equal(t, len(tt.expectCountMap), cnum)
			for k, expected := range tt.expectCountMap {
				actual, ok := countMap[k]
				assert.True(t, ok)
				assert.Equal(t, expected, actual)
			}
		})
	}
}
