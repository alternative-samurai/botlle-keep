package controllers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/src/infrastructures/model"
	"fmt"
)

func (r *mutationResolver) CreateCast(ctx context.Context, name string) (*model.IDResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CastList(ctx context.Context) (*model.CastListResult, error) {
	panic(fmt.Errorf("not implemented"))
}
