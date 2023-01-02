package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
)

const (
	DurationDB         = 1
	DurationKeyFormat  = "%.6f,%.6f:%.6f,%.6f"
	DurationExpiration = 0 // 期限なし
)

type FromToLatLng struct {
	From common.LatLng
	To   common.LatLng
}

func ToDurationKey(from, to common.LatLng) string {
	return fmt.Sprintf(DurationKeyFormat, from.Lat, from.Lng, to.Lat, to.Lng)
}

type DurationCacheClient struct {
	Client *redis.Client
}

func NewDurationCacheClient(host, port string) (*DurationCacheClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   DurationDB,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return &DurationCacheClient{
		Client: client,
	}, nil
}

func (c *DurationCacheClient) Get(
	ctx context.Context,
	fromTo FromToLatLng,
) (
	*time.Duration,
	bool,
) {
	key := ToDurationKey(fromTo.From, fromTo.To)
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, false
	}

	duration, err := time.ParseDuration(val)
	if err != nil {
		return nil, false
	}
	return &duration, true
}

func (c *DurationCacheClient) Set(
	ctx context.Context,
	fromTo FromToLatLng,
	duration time.Duration,
) error {
	key := ToDurationKey(fromTo.From, fromTo.To)
	err := c.Client.Set(ctx, key, duration.String(), DurationExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *DurationCacheClient) BulkGet(
	ctx context.Context,
	fromTos []FromToLatLng,
) (
	results []*time.Duration,
	notFoundIndex []int,
	err error,
) {
	keys := make([]string, len(fromTos))
	for i := range keys {
		keys[i] = ToDurationKey(fromTos[i].From, fromTos[i].To)
	}

	res, err := c.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, nil, err
	}

	results = make([]*time.Duration, len(res))
	for i := range res {
		v, ok := res[i].(string)
		if duration, err := time.ParseDuration(v); ok && err == nil {
			results[i] = &duration
		} else {
			notFoundIndex = append(notFoundIndex, i)
		}
	}

	return results, notFoundIndex, nil
}

func (c *DurationCacheClient) BulkSet(
	ctx context.Context,
	fromTos []FromToLatLng,
	durations []time.Duration,
) error {
	if len(fromTos) != len(durations) {
		return errors.New("fromTos and durations does not match length")
	}

	kv := make(map[string]string)
	for i, ft := range fromTos {
		key := ToDurationKey(ft.From, ft.To)
		kv[key] = durations[i].String()
	}

	err := c.Client.MSet(ctx, kv).Err()
	if err != nil {
		return err
	}

	return nil
}
