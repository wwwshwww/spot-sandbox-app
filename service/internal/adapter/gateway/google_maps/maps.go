package google_maps

import (
	"context"
	"time"

	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"googlemaps.github.io/maps"
)

const (
	distanceMatrixElementLimit = 25
)

type GoogleMapsClient struct {
	Client *maps.Client
}

func NewGoogleMapsClient(apikey string) (*GoogleMapsClient, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return nil, err
	}
	return &GoogleMapsClient{Client: c}, nil
}

func (c *GoogleMapsClient) GetDurationAndDistanceOneToMany(
	origin common.LatLng,
	destinations []common.LatLng,
) (
	durations []time.Duration, // without traffic duration
	distances []int, // meter
	err error,
) {
	durations = make([]time.Duration, len(destinations))
	distances = make([]int, len(destinations))

	// DistanceMatrixのorigin,destinationの上限を超えてリクエストしないよう分割する
	for s := 0; s <= len(destinations)/(distanceMatrixElementLimit+1); s++ {
		from := distanceMatrixElementLimit * s
		to := distanceMatrixElementLimit * (s + 1)
		if to > len(destinations) {
			to = len(destinations)
		}
		dests := make([]string, to-from)
		for i, d := range destinations[from:to] {
			dests[i] = d.String()
		}

		req := &maps.DistanceMatrixRequest{
			Origins:      []string{origin.String()},
			Destinations: dests,
			Mode:         maps.TravelModeDriving,
		}

		res, err := c.Client.DistanceMatrix(context.Background(), req)
		if err != nil {
			return nil, nil, err
		}

		for i, r := range res.Rows[0].Elements {
			durations[from+i] = r.Duration
			distances[from+i] = r.Distance.Meters
		}
	}

	return durations, distances, nil
}

func (c *GoogleMapsClient) GetDurationAndDistanceManyToOne(
	origins []common.LatLng,
	destination common.LatLng,
) (
	durations []time.Duration, // without traffic duration
	distances []int, // meter
	err error,
) {
	durations = make([]time.Duration, len(origins))
	distances = make([]int, len(origins))

	// DistanceMatrixのorigin,destinationの上限を超えてリクエストしないよう分割する
	for s := 0; s <= len(origins)/(distanceMatrixElementLimit+1); s++ {
		from := distanceMatrixElementLimit * s
		to := distanceMatrixElementLimit * (s + 1)
		if to > len(origins) {
			to = len(origins)
		}
		oris := make([]string, to-from)
		for i, o := range origins[from:to] {
			oris[i] = o.String()
		}

		req := &maps.DistanceMatrixRequest{
			Origins:      oris,
			Destinations: []string{destination.String()},
			Mode:         maps.TravelModeDriving,
		}

		res, err := c.Client.DistanceMatrix(context.Background(), req)
		if err != nil {
			return nil, nil, err
		}

		for i, r := range res.Rows {
			durations[from+i] = r.Elements[0].Duration
			distances[from+i] = r.Elements[0].Distance.Meters
		}
	}

	return durations, distances, nil
}

func (c *GoogleMapsClient) ReverseGeocode(lat, lng float64) (
	postalCode,
	addressRepresentation string,
	err error,
) {
	req := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		Language: "ja",
	}
	res, err := c.Client.ReverseGeocode(context.Background(), req)
	if err != nil {
		return
	}

	for _, c := range res[0].AddressComponents {
		for _, t := range c.Types {
			if t == "postal_code" {
				postalCode = c.LongName
				break
			}
		}
		if postalCode != "" {
			break
		}
	}
	addressRepresentation = res[0].FormattedAddress

	return postalCode, addressRepresentation, nil
}
