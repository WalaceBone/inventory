package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"inventory/graph/generated"
	"inventory/graph/model"
)

func (r *componentResolver) ID(ctx context.Context, obj *model.Component) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *componentResolver) Type(ctx context.Context, obj *model.Component) (model.ComponentType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *componentResolver) Quantity(ctx context.Context, obj *model.Component) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *recipeResolver) MandatoryComponent(ctx context.Context, obj *model.Recipe) ([]*model.Component, error) {
	panic(fmt.Errorf("not implemented"))
}

// Component returns generated.ComponentResolver implementation.
func (r *Resolver) Component() generated.ComponentResolver { return &componentResolver{r} }

// Recipe returns generated.RecipeResolver implementation.
func (r *Resolver) Recipe() generated.RecipeResolver { return &recipeResolver{r} }

type componentResolver struct{ *Resolver }
type recipeResolver struct{ *Resolver }
