package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/wwwwshwww/spot-sandbox/internal/common"
)

const (
	DistanceDB         = 0
	DistanceKeyFormat  = "%.6f,%.6f:%.6f,%.6f"
	DistanceExpiration = 0 // 期限なし
)

func ToDistanceKey(from, to common.LatLng) string {
	return fmt.Sprintf(DistanceKeyFormat, from.Lat, from.Lng, to.Lat, to.Lng)
}

type DistanceCacheClient struct {
	Client *redis.Client
}

func NewDistanceCacheClient(host, port string) (*DistanceCacheClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   DistanceDB,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return &DistanceCacheClient{
		Client: client,
	}, nil
}

func (c *DistanceCacheClient) Get(
	ctx context.Context,
	fromTo FromToLatLng,
) (
	*int,
	bool,
) {
	key := ToDistanceKey(fromTo.From, fromTo.To)
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, false
	}

	distance, err := strconv.Atoi(val)
	if err != nil {
		return nil, false
	}
	return &distance, true
}

func (c *DistanceCacheClient) Set(
	ctx context.Context,
	fromTo FromToLatLng,
	meter int,
) error {
	key := ToDistanceKey(fromTo.From, fromTo.To)
	err := c.Client.Set(ctx, key, strconv.Itoa(meter), DistanceExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *DistanceCacheClient) BulkGet(
	ctx context.Context,
	fromTos []FromToLatLng,
) (
	results []*int,
	notFoundIndex []int,
	err error,
) {
	keys := make([]string, len(fromTos))
	for i := range keys {
		keys[i] = ToDistanceKey(fromTos[i].From, fromTos[i].To)
	}

	res, err := c.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, nil, err
	}

	results = make([]*int, len(res))
	for i := range res {
		v, ok := res[i].(string)
		if distance, err := strconv.Atoi(v); ok && err == nil {
			results[i] = &distance
		} else {
			notFoundIndex = append(notFoundIndex, i)
		}
	}

	return results, notFoundIndex, nil
}

func (c *DistanceCacheClient) BulkSet(
	ctx context.Context,
	fromTos []FromToLatLng,
	meters []int,
) error {
	if len(fromTos) != len(meters) {
		return errors.New("fromTos and merters does not match length")
	}

	kv := make(map[string]string)
	for i, ft := range fromTos {
		key := ToDistanceKey(ft.From, ft.To)
		kv[key] = strconv.Itoa(meters[i])
	}

	err := c.Client.MSet(ctx, kv).Err()
	if err != nil {
		return err
	}

	return nil
}
