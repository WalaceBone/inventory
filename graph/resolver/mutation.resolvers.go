package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"inventory/graph/generated"
	"inventory/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateIngredient(ctx context.Context, input model.NewIngredient) (*model.Ingredient, error) {
	ingredient := &model.Ingredient{
		ID:           fmt.Sprintf("T%d", rand.Int()),
		Name:         input.Name,
		CurrentStock: input.CurrentStock,
		MinimumStock: input.MinimumStock,
	}
	r.ingredients = append(r.ingredients, ingredient)
	return ingredient, nil
}

func (r *mutationResolver) UpdateingredientStock(ctx context.Context, quantity int, ingredientID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRecipe(ctx context.Context, input model.NewRecipe) (*model.Recipe, error) {
	recipe := &model.Recipe{
		ID:                 fmt.Sprintf("T%d", rand.Int()),
		Name:               input.Name,
		CurrentStock:       input.CurrentStock,
		MinimumStock:       input.MinimumStock,
		MandatoryComponent: model.Component{},
	}
	r.recipes = append(r.recipes, recipe)
	return recipe, nil
}

func (r *mutationResolver) UpdateRecipeStock(ctx context.Context, quantity int, recipeID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRecipeComponentQuantity(ctx context.Context, quantity int, recipeID string, componentID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddRecipeComponent(ctx context.Context, recipeID string, component model.ComponentInput) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRecipeComponent(ctx context.Context, recipeID string, componentID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
