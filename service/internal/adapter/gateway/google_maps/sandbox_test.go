package google_maps_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/wwwwshwww/spot-sandbox/internal/config"
	"googlemaps.github.io/maps"
)

func TestAPI(t *testing.T) {
	config.Configure()

	c, err := maps.NewClient(maps.WithAPIKey(config.GoogleMapsAPIKey))
	if err != nil {
		panic(err)
	}

	r := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: 35.562479,
			Lng: 139.716073,
		},
		Language: "ja",
	}

	res, err := c.ReverseGeocode(context.Background(), r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}

func TestDistanceMatrix(t *testing.T) {
	config.Configure()

	c, err := maps.NewClient(maps.WithAPIKey(config.GoogleMapsAPIKey))
	if err != nil {
		panic(err)
	}

	r := &maps.DistanceMatrixRequest{
		Origins: []string{
			fmt.Sprintf("%f,%f", 35.562479, 139.716073),
		},
		Destinations: []string{
			fmt.Sprintf("%f,%f", 35.549497, 139.714922),
			fmt.Sprintf("%f,%f", 35.540531, 139.707582),
		},
		Mode:     maps.TravelModeDriving,
		Language: "Japanese",
	}

	_ = maps.DistanceMatrixResponse{
		Rows: []maps.DistanceMatrixElementsRow{
			{
				Elements: []*maps.DistanceMatrixElement{
					{},
				},
			},
		},
	}

	res, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}
