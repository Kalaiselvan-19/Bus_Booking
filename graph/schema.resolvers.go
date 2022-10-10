package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Bus_Booking/graph/generated"
	"Bus_Booking/graph/model"
	"Bus_Booking/service"
	"context"
)

// Login is the resolver for the login field.
func (r *authResolver) Login(ctx context.Context, obj *model.Auth, email string, password string) (interface{}, error) {
	return service.UserLogin(ctx, email, password)
}

// Register is the resolver for the register field.
func (r *authResolver) Register(ctx context.Context, obj *model.Auth, input model.NewUser) (interface{}, error) {
	return service.UserRegister(ctx, input)
}

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*model.Auth, error) {
	return &model.Auth{}, nil
}

// Getuser is the resolver for the getuser field.
func (r *queryResolver) Getuser(ctx context.Context, input model.GetUser) (*model.User, error) {
	return service.UserGetByEmail(ctx, input.Email)
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

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "Success", nil
}

// Auth returns generated.AuthResolver implementation.
func (r *Resolver) Auth() generated.AuthResolver { return &authResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
