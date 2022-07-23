package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wwwwshwww/spot-sandbox/entity"
	"github.com/wwwwshwww/spot-sandbox/graph/generated"
	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/loader"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	record := entity.User{
		Name: input.Name,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewUserFromEntity(&record)

	return res, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	userID, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}
	record := entity.Todo{
		Text:   input.Text,
		UserID: uint(userID),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewTodoFromEntity(&record)
	return res, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	idn, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var u entity.User
	if err := r.DB.Find(&u, idn).Error; err != nil {
		return nil, err
	}

	return &model.User{
		ID:   fmt.Sprintf("%d", u.ID),
		Name: u.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var records []entity.Todo
	if err := r.DB.Find(&records).Error; err != nil {
		return nil, err
	}

	todos := []*model.Todo{}
	for _, record := range records {
		todos = append(todos, model.NewTodoFromEntity(&record))
	}
	return todos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	// userID, err := strconv.Atoi(obj.UserID)
	// if err != nil {
	// 	return nil, err
	// }
	// var record entity.User
	// if err := r.DB.Debug().Find(&record, userID).Error; err != nil {
	// 	return nil, err
	// }
	// return model.NewUserFromEntity(&record), nil
	return loader.LoadUser(ctx, obj.UserID)
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	userID, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}
	var records []entity.Todo
	if err := r.DB.Where("user_id", userID).Find(&records).Error; err != nil {
		return nil, err
	}

	todos := []*model.Todo{}
	for _, record := range records {
		todos = append(todos, model.NewTodoFromEntity(&record))
	}

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
