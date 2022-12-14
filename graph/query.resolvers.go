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
	subject, err := r.sbU.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` func of SubjectUsecase: %w", err)
	}
	s := model.Subject(*subject)
	return &s, nil
}

// Subjects is the resolver for the subjects field.
func (r *queryResolver) Subjects(ctx context.Context, title string, faculty string, academicField string) ([]*model.Subject, error) {
	if title == "" && faculty == "" && academicField == "" {
		return nil, fmt.Errorf("at least one of the parameters must be specified")
	}
	subjects, err := r.sbU.GetBySearchParameter(title, faculty, academicField)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetBySearchParameter` func of SubjectUsecase: %w", err)
	}
	ss := make([]*model.Subject, len(subjects))
	for i, subject := range subjects {
		s := model.Subject(*subject)
		ss[i] = &s
	}
	return ss, nil
}

// AcademicFields is the resolver for the academicFields field.
func (r *queryResolver) AcademicFields(ctx context.Context) ([]*model.AcademicField, error) {
	academicFieldNames, err := r.afU.Get()
	if err != nil {
		return nil, fmt.Errorf("failed on executing `Get` func of AcademicFieldUsecase: %w", err)
	}
	academicFields := make([]*model.AcademicField, len(academicFieldNames))
	for i, name := range academicFieldNames {
		academicFields[i] = &model.AcademicField{
			Name: name,
		}
	}
	return academicFields, nil
}

// RandomSubjects is the resolver for the randomSubjects field.
func (r *queryResolver) RandomSubjects(ctx context.Context, category string, series string, academicField string, numSubjects int) ([]*model.Subject, error) {
	subjects, err := r.sbU.GetByRandom(category, series, academicField, numSubjects)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByRandom` func of SubjectUsecase: %w", err)
	}

	ss := make([]*model.Subject, len(subjects))
	for i, subject := range subjects {
		s := model.Subject(*subject)
		ss[i] = &s
	}
	return ss, nil
}

// SubjectsWithSpecifiedVideos is the resolver for the subjectsWithSpecifiedVideos field.
func (r *queryResolver) SubjectsWithSpecifiedVideos(ctx context.Context, title string, faculty string) ([]*model.SubjectWithSpecifiedVideos, error) {
	if title == "" && faculty == "" {
		return nil, fmt.Errorf("at least one of the parameters must be specified")
	}
	subjectWithSpecifiedVideoDTOs, err := r.sbU.GetByVideoSearchParameter(title, faculty)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByVideoSearchParameter` func of SubjectUsecase: %w", err)
	}
	svs := make([]*model.SubjectWithSpecifiedVideos, len(subjectWithSpecifiedVideoDTOs))
	for i, sv := range subjectWithSpecifiedVideoDTOs {
		svs[i] = model.NewSubjectWithSpecifiedVideos(sv)
	}
	return svs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
