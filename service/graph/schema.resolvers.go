package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/wwwwshwww/spot-sandbox/graph/model"
	dbscan_profile_graph "github.com/wwwwshwww/spot-sandbox/internal/adapter/inbound/dbscan_profile/graph"
	spot_graph "github.com/wwwwshwww/spot-sandbox/internal/adapter/inbound/spot/graph"
	dbscan_profile_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/dbscan_profile/mysql"
	spot_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/spot/spot/mysql"
	spot_finder_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/spot/spot_finder/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot_finder"
	"github.com/wwwwshwww/spot-sandbox/internal/usecase/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/usecase/spot"
)

// DbscanProfile is the resolver for the dbscanProfile field.
func (r *clusterElementResolver) DbscanProfile(ctx context.Context, obj *model.ClusterElement) (*model.DbscanProfile, error) {
	panic(fmt.Errorf("not implemented: DbscanProfile - dbscanProfile"))
}

// SpotsProfile is the resolver for the spotsProfile field.
func (r *clusterElementResolver) SpotsProfile(ctx context.Context, obj *model.ClusterElement) (*model.SpotsProfile, error) {
	panic(fmt.Errorf("not implemented: SpotsProfile - spotsProfile"))
}

// Spot is the resolver for the spot field.
func (r *clusterElementResolver) Spot(ctx context.Context, obj *model.ClusterElement) (*model.Spot, error) {
	panic(fmt.Errorf("not implemented: Spot - spot"))
}

// CreateDbscanProfile is the resolver for the createDbscanProfile field.
func (r *mutationResolver) CreateDbscanProfile(ctx context.Context, input model.NewDbscanProfile) (*model.DbscanProfile, error) {
	dpr := dbscan_profile_mysql.New(r.DB)
	dpuc := dbscan_profile.New(dpr)

	i, err := dpr.NextIdentifier()
	if err != nil {
		return nil, err
	}
	p := dbscan_profile_graph.UnmarshalPreferences(input)
	if err := dpuc.Save(i, p); err != nil {
		return nil, err
	}
	dp, err := dpuc.Get(i)
	if err != nil {
		return nil, err
	}
	mdp := dbscan_profile_graph.Marshal(dp)
	return &mdp, nil
}

// Spots is the resolver for the spots field.
func (r *queryResolver) Spots(ctx context.Context) ([]*model.Spot, error) {
	sr := spot_mysql.New(r.DB)
	sf := spot_finder_mysql.New(r.DB)
	suc := spot.New(sr, sf, r.GMC)

	sis, err := suc.ListAllSpots(spot_finder.FilteringOptions{})
	if err != nil {
		return nil, err
	}
	sMap, err := suc.BulkGet(sis)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Spot, len(sis))
	for i, v := range sis {
		result[i] = spot_graph.Marshal(sMap[v])
	}
	return result, nil
}

// Spots is the resolver for the spots field.
func (r *spotsProfileResolver) Spots(ctx context.Context, obj *model.SpotsProfile) ([]*model.Spot, error) {
	panic(fmt.Errorf("not implemented: Spots - spots"))
}

// ClusterElement returns ClusterElementResolver implementation.
func (r *Resolver) ClusterElement() ClusterElementResolver { return &clusterElementResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// SpotsProfile returns SpotsProfileResolver implementation.
func (r *Resolver) SpotsProfile() SpotsProfileResolver { return &spotsProfileResolver{r} }

type clusterElementResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type spotsProfileResolver struct{ *Resolver }
