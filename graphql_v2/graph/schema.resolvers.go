package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql_v2/graph/generated"
	"graphql_v2/graph/internal/provincestore"
	"graphql_v2/graph/model"
)

func (r *mutationResolver) CreateProvince(ctx context.Context, input model.NewProvince) (*model.Province, error) {
	modifiers := make([]*model.Modifier, 0, len(input.Modifiers))
	for _, a := range input.Modifiers {
		modifiers = append(modifiers, (*model.Modifier)(a))
	}
	id := r.Store.CreateProvince(input.ID, input.Name, modifiers)
	province, err := r.Store.GetProvince(id)
	return province, err
}

func (r *mutationResolver) DeleteProvince(ctx context.Context, id string) (*bool, error) {
	return nil, r.Store.DeleteProvince(id)
}

func (r *mutationResolver) DeleteAllProvinces(ctx context.Context) (*bool, error) {
	if r == nil {
		r = provincestore.New()
	}
	return nil, r.Store.DeleteAllProvinces()
}

func (r *queryResolver) GetAllProvinces(ctx context.Context) ([]*model.Province, error) {
	return r.Store.GetAllProvinces(), nil
}

func (r *queryResolver) GetProvince(ctx context.Context, id string) (*model.Province, error) {
	return r.Store.GetProvince(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
