package graph

import (
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/google_maps"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB  *gorm.DB
	GMC *google_maps.GoogleMapsClient
}
