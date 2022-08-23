package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kafugen/ocwcentral/graph/generated"
	"github.com/kafugen/ocwcentral/graph/model"
)

// Videos is the resolver for the videos field.
func (r *subjectResolver) Videos(ctx context.Context, obj *model.Subject) ([]*model.Video, error) {
	panic(fmt.Errorf("not implemented: Videos - videos"))
}

// Resources is the resolver for the resources field.
func (r *subjectResolver) Resources(ctx context.Context, obj *model.Subject) ([]*model.Resource, error) {
	panic(fmt.Errorf("not implemented: Resources - resources"))
}

// RelatedSubjects is the resolver for the relatedSubjects field.
func (r *subjectResolver) RelatedSubjects(ctx context.Context, obj *model.Subject) ([]*model.RelatedSubject, error) {
	panic(fmt.Errorf("not implemented: RelatedSubjects - relatedSubjects"))
}

// Syllabus is the resolver for the syllabus field.
func (r *subjectResolver) Syllabus(ctx context.Context, obj *model.Subject) (*model.Syllabus, error) {
	panic(fmt.Errorf("not implemented: Syllabus - syllabus"))
}

// Subject returns generated.SubjectResolver implementation.
func (r *Resolver) Subject() generated.SubjectResolver { return &subjectResolver{r} }

type subjectResolver struct{ *Resolver }
