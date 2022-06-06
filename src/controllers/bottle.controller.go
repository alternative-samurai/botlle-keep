package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	generated1 "example/src/infrastructures/generated"
	model1 "example/src/infrastructures/model"
	"fmt"
)

func (r *mutationResolver) CreateBottle(ctx context.Context, name string) (*model1.IDResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BottleList(ctx context.Context) (*model1.BottleListResult, error) {
	return &model1.BottleListResult{
		List: []*model1.Bottle{
			{BottleID: "aaaa", BottleName: "bbb"},
		},
	}, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
