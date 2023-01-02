package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/cache"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/config"
)

func TestLatLngCache(t *testing.T) {
	config.Configure()
	ctx := context.Background()

	clientDist, err := cache.NewDistanceCacheClient(config.Redis.Host, config.Redis.Port)
	assert.NoError(t, err)
	clientDura, err := cache.NewDurationCacheClient(config.Redis.Host, config.Redis.Port)
	assert.NoError(t, err)

	count := 5
	keys := make([]cache.FromToLatLng, count)
	meters := make([]int, count)
	durations := make([]time.Duration, count)
	for i := 0; i < count; i++ {
		keys[i] = cache.FromToLatLng{
			From: common.LatLng{Lat: 0.0, Lng: 0.0},
			To:   common.LatLng{Lat: 0.1 * float64(i), Lng: 0.1 * float64(i)},
		}
		meters[i] = i * 10
		durations[i] = time.Second * time.Duration(i*10)
	}

	err = clientDist.BulkSet(ctx, keys, meters)
	assert.NoError(t, err)
	err = clientDura.BulkSet(ctx, keys, durations)
	assert.NoError(t, err)

	// Distance: BulkGetテスト
	resDist, nfDist, err := clientDist.BulkGet(ctx, keys)
	assert.NoError(t, err)
	assert.ElementsMatch(
		t,
		common.Map(func(n int) *int { return &n }, meters),
		resDist,
	)
	assert.Equal(t, 0, len(nfDist))

	// Distance: NotFountが含まれるBulkGetテスト
	keysAmountNF := append(keys, cache.FromToLatLng{
		From: common.LatLng{Lat: 99.0, Lng: 99.0},
		To:   common.LatLng{Lat: 100.0, Lng: 100.0},
	})
	resDist, nfDist, err = clientDist.BulkGet(ctx, keysAmountNF)
	assert.NoError(t, err)
	assert.ElementsMatch(
		t,
		append(
			common.Map(func(n int) *int { return &n }, meters),
			nil,
		),
		resDist,
	)
	assert.Equal(t, 1, len(nfDist))

	// Duration: BulkGetテスト
	resDura, nfDura, err := clientDura.BulkGet(ctx, keys)
	assert.NoError(t, err)
	assert.ElementsMatch(
		t,
		common.Map(func(n time.Duration) *time.Duration { return &n }, durations),
		resDura,
	)
	assert.Equal(t, 0, len(nfDura))

	// Distance: NotFountが含まれるBulkGetテスト
	resDura, nfDura, err = clientDura.BulkGet(ctx, keysAmountNF)
	assert.NoError(t, err)
	assert.ElementsMatch(
		t,
		append(
			common.Map(func(n time.Duration) *time.Duration { return &n }, durations),
			nil,
		),
		resDura,
	)
	assert.Equal(t, 1, len(nfDura))
}
