package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"inventory/graph/generated"
	"inventory/graph/model"
)

func (r *queryResolver) Ingredients(ctx context.Context) ([]*model.Ingredient, error) {
	return r.ingredients, nil
}

func (r *queryResolver) Recipes(ctx context.Context) ([]*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
