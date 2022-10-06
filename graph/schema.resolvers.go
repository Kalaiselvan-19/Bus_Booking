package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Bus_Booking/graph/generated"
	"Bus_Booking/graph/model"
	"context"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber}
	_, err := r.DB.Model(&user).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new user: %v", err)
	}
	return &user, nil
}

// Getuser is the resolver for the getuser field.
func (r *queryResolver) Getuser(ctx context.Context, input model.GetUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Select(input.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Alluser is the resolver for the alluser field.
func (r *queryResolver) Alluser(ctx context.Context) ([]*model.User, error) {
	var user []*model.User

	err := r.DB.Model(&user).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
