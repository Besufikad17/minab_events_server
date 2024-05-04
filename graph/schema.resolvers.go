package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/Besufikad17/minab_events/graph/hasura/actions"
	models "github.com/Besufikad17/minab_events/graph/hasura/models"
	"github.com/Besufikad17/minab_events/graph/model"
	"github.com/Besufikad17/minab_events/graph/utils/helpers"
)

// Register is the resolver for the Register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	hashedPassword, err := helpers.Hash(input.Password)

	if err != nil {
		return nil, err
	}

	createUserInput := models.RegisterArgs{
		First_name:   input.FirstName,
		Last_name:    input.LastName,
		Email:        input.Email,
		Phone_number: input.PhoneNumber,
		Password:     hashedPassword,
	}

	result, err := actions.Register(createUserInput)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:          *&result.Id,
		FirstName:   *result.First_name,
		LastName:    *result.Last_name,
		Email:       *result.Email,
		PhoneNumber: *result.Phone_number,
	}
	token, err := helpers.CreateToken(*user)
	user.Token = token

	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
