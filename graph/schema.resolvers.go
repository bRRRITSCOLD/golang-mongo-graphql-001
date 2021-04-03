package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-mongo-graphql-001/graph/generated"
	"golang-mongo-graphql-001/graph/model"
)

func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
