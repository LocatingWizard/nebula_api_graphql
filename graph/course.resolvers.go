package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/LocatingWizard/nebula_api_graphql/graph/model"
)

// Sections is the resolver for the sections field.
func (r *courseResolver) Sections(ctx context.Context, obj *model.Course) ([]string, error) {
	panic(fmt.Errorf("not implemented: Sections - sections"))
}

// Course returns CourseResolver implementation.
func (r *Resolver) Course() CourseResolver { return &courseResolver{r} }

type courseResolver struct{ *Resolver }
