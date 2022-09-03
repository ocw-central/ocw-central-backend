package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kafugen/ocwcentral/graph/generated"
	"github.com/kafugen/ocwcentral/graph/model"
	"github.com/kafugen/ocwcentral/utils"
)

// Subject is the resolver for the subject field.
func (r *queryResolver) Subject(ctx context.Context, id string) (*model.Subject, error) {
	subject, err := r.sbU.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` func of SubjectUsecase: %w", err)
	}
	s := model.Subject(*subject)
	return &s, nil
}

// Subjects is the resolver for the subjects field.
func (r *queryResolver) Subjects(ctx context.Context, title *string, faculty *string, academicField *string) ([]*model.Subject, error) {
	subjectSearchParameter := utils.SubjectSearchParameter{
		Title:         utils.ConvertNilToZeroValue(title),
		Faculty:       utils.ConvertNilToZeroValue(faculty),
		AcademicField: utils.ConvertNilToZeroValue(academicField),
	}
	subjects, err := r.sbU.GetBySearchParameter(subjectSearchParameter)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByTitleAndDepartment` func of SubjectUsecase: %w", err)
	}
	ss := make([]*model.Subject, len(subjects))
	for i, subject := range subjects {
		s := model.Subject(*subject)
		ss[i] = &s
	}
	return ss, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
