package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kafugen/ocwcentral/graph/generated"
	"github.com/kafugen/ocwcentral/graph/model"
)

// Subject is the resolver for the subject field.
func (r *queryResolver) Subject(ctx context.Context, id string) (*model.Subject, error) {
	panic(fmt.Errorf("not implemented: Subject - Subject"))
}

// Subjects is the resolver for the subjects field.
func (r *queryResolver) Subjects(ctx context.Context, title *string, department *string) ([]*model.Subject, error) {
	panic(fmt.Errorf("not implemented: Subjects - subjects"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
