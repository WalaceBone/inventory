package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"inventory/graph/model"
)

func (r *ingredientResolver) CurrentStock(ctx context.Context, obj *model.Ingredient) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ingredientResolver) MinimumStock(ctx context.Context, obj *model.Ingredient) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ingredient returns generated.IngredientResolver implementation.
func (r *Resolver) Ingredient() *ingredientResolver { return &ingredientResolver{r} }

type ingredientResolver struct{ *Resolver }
