package controllers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/src/domains/bottle"
	"example/src/infrastructures/model"
	"fmt"
)

func (r *mutationResolver) CreateBottle(ctx context.Context, name string) (*model.IDResult, error) {
	fmt.Print(ctx.Value("reqDate"))

	return &model.IDResult{
		ID: bottle.CreateBottle(ctx, name),
	}, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BottleList(ctx context.Context) (*model.BottleListResult, error) {
	panic(fmt.Errorf("not implemented"))
}
